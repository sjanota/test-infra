package tests

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"
)

var tests = []struct {
	path              string
	image             string
	additionalOptions []buildjob.Option
}{
	{path: "end-to-end/backup-restore-test", image: tester.ImageGolangBuildpack1_11},
	//{path: "end-to-end/external-solution-integration", image: tester.ImageGolangBuildpack1_11},
	{path: "end-to-end/kubeless-integration", image: tester.ImageGolangBuildpack1_11},
	{path: "end-to-end/upgrade", image: tester.ImageGolangBuildpack1_11, additionalOptions: []buildjob.Option{
		buildjob.RunIfChanged("^tests/end-to-end/upgrade/[^chart]", "tests/end-to-end/upgrade/fix"),
	}},
}

func TestTestJobs(t *testing.T) {
	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			opts := []buildjob.Option{
				buildjob.Test(test.path, test.image),
				buildjob.KymaRepo(),
				buildjob.AllReleases(),
				buildjob.RepositoryRoot("./../../../../../../"),
			}
			opts = append(opts, test.additionalOptions...)
			buildjob.NewSuite(opts...).Run(t)
		})
	}
}
