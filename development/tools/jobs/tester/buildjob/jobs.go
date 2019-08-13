package buildjob

import (
	"fmt"
	. "github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/test-infra/prow/config"
	"path"
	"strings"
	"testing"
)

type Suite struct {
	Path              string
	Repository        string
	Image             string
	Releases          []*SupportedRelease
	RunIfChanged      string
	RunIfChangedCheck string
}

type Option func(suite *Suite)

func NewSuite(opts ...Option) *Suite {
	suite := &Suite{
		Releases: GetAllKymaReleases(),
	}
	for _, opt := range opts {
		opt(suite)
	}
	setDefaults(suite)
	return suite
}

func Component(name, image string) Option {
	return func(suite *Suite) {
		suite.Path = fmt.Sprintf("components/%s", name)
		suite.Image = image
	}
}

func Test(name, image string) Option {
	return func(suite *Suite) {
		suite.Path = fmt.Sprintf("tests/%s", name)
		suite.Image = image
	}
}

func KymaRepo() Option {
	return func(suite *Suite) {
		suite.Repository = "github.com/kyma-project/kyma"
	}
}

func setDefaults(s *Suite) {
	if s.RunIfChanged == "" {
		s.RunIfChanged = fmt.Sprintf("^%s/", s.Path)
	}
	if s.RunIfChangedCheck == "" {
		s.RunIfChangedCheck = fmt.Sprintf("%s/fix", s.Path)
	}
}

func (s *Suite) componentName() string {
	return path.Base(s.Path)
}

func (s *Suite) repositoryName() string {
	return path.Base(s.Repository)
}

func (s *Suite) repositorySectionKey() string {
	return strings.Replace(s.Repository, "github.com/", "", 1)
}

func (s *Suite) moduleName() string {
	return fmt.Sprintf("%s-%s", s.repositoryName(), strings.Replace(s.Path, "/", "-", -1))
}

func (s *Suite) jobConfigPath() string {
	return fmt.Sprintf("./../../../../prow/jobs/%s/%s/%s.yaml", s.repositoryName(), s.Path, s.componentName())
}

func (s *Suite) jobName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, s.moduleName())
}

func (s *Suite) workingdirectory() string {
	return fmt.Sprintf("/home/prow/go/src/%s/%s", s.Repository, s.Path)
}

func (s *Suite) Run(t *testing.T) {
	jobConfig, err := ReadJobConfig(s.jobConfigPath())
	require.NoError(t, err)

	expectedNumberOfPresubmits := len(s.Releases) + 1
	require.Len(t, jobConfig.Presubmits, 1)
	require.Len(t, jobConfig.Presubmits[s.repositorySectionKey()], expectedNumberOfPresubmits)

	require.Len(t, jobConfig.Postsubmits, 1)
	require.Len(t, jobConfig.Postsubmits[s.repositorySectionKey()], 1)

	require.Empty(t, jobConfig.Periodics)

	t.Run("pre-master", s.preMasterTest(jobConfig))
	t.Run("post-master", s.postMasterTest(jobConfig))
	t.Run("release", s.preReleaseTest(jobConfig))
}

func (s *Suite) preMasterTest(jobConfig config.JobConfig) func(t *testing.T) {
	return func(t *testing.T) {
		actualPresubmit := FindPresubmitJobByName(
			jobConfig.Presubmits[s.repositorySectionKey()],
			s.jobName("pre-master"),
			"master",
		)
		require.NotNil(t, actualPresubmit)

		assert.Equal(t, []string{"^master$"}, actualPresubmit.Branches)
		assert.Equal(t, 10, actualPresubmit.MaxConcurrency)
		assert.False(t, actualPresubmit.SkipReport)
		assert.True(t, actualPresubmit.Decorate)
		assert.Equal(t, s.Repository, actualPresubmit.PathAlias)
		AssertThatHasExtraRefTestInfra(t, actualPresubmit.JobBase.UtilityConfig, "master")
		AssertThatHasPresets(t, actualPresubmit.JobBase, PresetDindEnabled, PresetDockerPushRepo, PresetGcrPush, PresetBuildPr)
		AssertThatJobRunIfChanged(t, *actualPresubmit, s.RunIfChangedCheck)
		assert.Equal(t, s.RunIfChanged, actualPresubmit.RunIfChanged)
		AssertThatExecGolangBuildpack(t, actualPresubmit.JobBase, s.Image, s.workingdirectory())
	}
}

func (s *Suite) postMasterTest(jobConfig config.JobConfig) func(t *testing.T) {
	return func(t *testing.T) {
		actualPostsubmit := FindPostsubmitJobByName(
			jobConfig.Postsubmits[s.repositorySectionKey()],
			s.jobName("post-master"),
			"master",
		)
		require.NotNil(t, actualPostsubmit)

		assert.Equal(t, []string{"^master$"}, actualPostsubmit.Branches)
		assert.Equal(t, 10, actualPostsubmit.MaxConcurrency)
		assert.True(t, actualPostsubmit.Decorate)
		assert.Equal(t, s.Repository, actualPostsubmit.PathAlias)
		AssertThatHasExtraRefTestInfra(t, actualPostsubmit.JobBase.UtilityConfig, "master")
		AssertThatHasPresets(t, actualPostsubmit.JobBase, PresetDindEnabled, PresetDockerPushRepo, PresetGcrPush, PresetBuildMaster)
		assert.Equal(t, s.RunIfChanged, actualPostsubmit.RunIfChanged)
		AssertThatExecGolangBuildpack(t, actualPostsubmit.JobBase, s.Image, s.workingdirectory())
	}
}

func (s *Suite) preReleaseTest(jobConfig config.JobConfig) func(t *testing.T) {
	return func(t *testing.T) {
		for _, currentRelease := range s.Releases {
			t.Run(currentRelease.String(), func(t *testing.T) {
				actualPresubmit := FindPresubmitJobByName(
					jobConfig.Presubmits[s.repositorySectionKey()],
					GetReleaseJobName(s.moduleName(), currentRelease),
					currentRelease.Branch(),
				)
				require.NotNil(t, actualPresubmit)

				assert.Equal(t, []string{currentRelease.Branch()}, actualPresubmit.Branches)
				assert.False(t, actualPresubmit.SkipReport)
				assert.True(t, actualPresubmit.Decorate)
				assert.Equal(t, s.Repository, actualPresubmit.PathAlias)
				AssertThatHasExtraRefTestInfra(t, actualPresubmit.JobBase.UtilityConfig, currentRelease.Branch())
				AssertThatHasPresets(t, actualPresubmit.JobBase, PresetDindEnabled, PresetDockerPushRepo, PresetGcrPush, PresetBuildRelease)
				AssertThatJobRunIfChanged(t, *actualPresubmit, s.RunIfChangedCheck)
				assert.True(t, actualPresubmit.AlwaysRun)
				AssertThatExecGolangBuildpack(t, actualPresubmit.JobBase, s.Image, s.workingdirectory())
			})
		}
	}
}
