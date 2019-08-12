package kyma_test

import (
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEtcdTlsSetupJobsPresubmit(t *testing.T) {
	// WHEN
	jobConfig, err := tester.ReadJobConfig("./../../../../prow/jobs/kyma/components/etcd-tls-setup-job/etcd-tls-setup-job.yaml")
	// THEN
	require.NoError(t, err)

	actualPresubmit := tester.FindPresubmitJobByName(jobConfig.Presubmits["kyma-project/kyma"], "pre-master-kyma-components-etcd-tls-setup-job", "master")
	require.NotNil(t, actualPresubmit)
	assert.Equal(t, []string{"^master$"}, actualPresubmit.Branches)
	assert.Equal(t, "github.com/kyma-project/kyma", actualPresubmit.PathAlias)
	tester.AssertThatHasExtraRefTestInfra(t, actualPresubmit.JobBase.UtilityConfig, "master")
	tester.AssertThatHasPresets(t, actualPresubmit.JobBase, tester.PresetDindEnabled, tester.PresetDockerPushRepo, tester.PresetGcrPush, tester.PresetBuildPr)
	assert.Equal(t, "^components/etcd-tls-setup-job/", actualPresubmit.RunIfChanged)
	assert.Equal(t, tester.ImageGolangBuildpack1_11, actualPresubmit.Spec.Containers[0].Image)
	assert.Equal(t, []string{"/home/prow/go/src/github.com/kyma-project/test-infra/prow/scripts/build.sh"}, actualPresubmit.Spec.Containers[0].Command)
	assert.Equal(t, []string{"/home/prow/go/src/github.com/kyma-project/kyma/components/etcd-tls-setup-job"}, actualPresubmit.Spec.Containers[0].Args)
}

func TestEtcdTlsSetupPostsubmit(t *testing.T) {
	// WHEN
	jobConfig, err := tester.ReadJobConfig("./../../../../prow/jobs/kyma/components/etcd-tls-setup-job/etcd-tls-setup-job.yaml")
	// THEN
	require.NoError(t, err)

	actualPost := tester.FindPostsubmitJobByName(jobConfig.Postsubmits["kyma-project/kyma"], "post-master-kyma-components-etcd-tls-setup-job", "master")
	require.NotNil(t, actualPost)

	assert.Equal(t, []string{"^master$"}, actualPost.Branches)
	assert.Equal(t, "github.com/kyma-project/kyma", actualPost.PathAlias)
	tester.AssertThatHasExtraRefTestInfra(t, actualPost.JobBase.UtilityConfig, "master")
	tester.AssertThatHasPresets(t, actualPost.JobBase, tester.PresetDindEnabled, tester.PresetDockerPushRepo, tester.PresetGcrPush, tester.PresetBuildMaster)
	assert.Equal(t, "^components/etcd-tls-setup-job/", actualPost.RunIfChanged)
	assert.Equal(t, tester.ImageGolangBuildpack1_11, actualPost.Spec.Containers[0].Image)
	assert.Equal(t, []string{"/home/prow/go/src/github.com/kyma-project/test-infra/prow/scripts/build.sh"}, actualPost.Spec.Containers[0].Command)
	assert.Equal(t, []string{"/home/prow/go/src/github.com/kyma-project/kyma/components/etcd-tls-setup-job"}, actualPost.Spec.Containers[0].Args)
}

func TestEtcdTlsSetupReleases(t *testing.T) {
	// WHEN
	for _, currentRelease := range tester.GetAllKymaReleaseBranches() {
		t.Run(currentRelease, func(t *testing.T) {
			expectedImage := tester.ImageGolangBuildpack1_11
			jobConfig, err := tester.ReadJobConfig("./../../../../prow/jobs/kyma/components/etcd-tls-setup-job/etcd-tls-setup-job.yaml")
			// THEN
			require.NoError(t, err)
			actualPresubmit := tester.FindPresubmitJobByName(jobConfig.Presubmits["kyma-project/kyma"], tester.GetReleaseJobName("kyma-components-etcd-tls-setup-job", currentRelease), currentRelease)
			require.NotNil(t, actualPresubmit)
			assert.False(t, actualPresubmit.SkipReport)
			assert.True(t, actualPresubmit.Decorate)
			assert.Equal(t, "github.com/kyma-project/kyma", actualPresubmit.PathAlias)
			tester.AssertThatHasExtraRefTestInfra(t, actualPresubmit.JobBase.UtilityConfig, currentRelease)
			tester.AssertThatHasPresets(t, actualPresubmit.JobBase, tester.PresetDindEnabled, tester.PresetDockerPushRepo, tester.PresetGcrPush, tester.PresetBuildRelease)
			assert.True(t, actualPresubmit.AlwaysRun)
			tester.AssertThatExecGolangBuildpack(t, actualPresubmit.JobBase, expectedImage, "/home/prow/go/src/github.com/kyma-project/kyma/components/etcd-tls-setup-job")
		})
	}
}
