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
	{path: "integration/api-controller", image: tester.ImageGolangBuildpack1_12,
		additionalOptions:[]buildjob.Option{
			buildjob.JobFileSuffix("tests"),
		},
	},
	{path: "integration/apiserver-proxy", image: tester.ImageGolangBuildpack1_12,
		additionalOptions:[]buildjob.Option{
			buildjob.JobFileSuffix("tests"),
		},
	},
	{path: "integration/cluster-users", image: tester.ImageBootstrapLatest},
	{path: "integration/dex", image: tester.ImageGolangBuildpack1_12},
	{path: "integration/event-service", image: tester.ImageGolangBuildpack1_11,
		additionalOptions:[]buildjob.Option{
			buildjob.JobFileSuffix("tests"),
		},
	},
	{path: "integration/logging", image: tester.ImageGolangBuildpackLatest, additionalOptions: []buildjob.Option{
		buildjob.Since(tester.Release14),
	}},
	{path: "integration/monitoring", image: tester.ImageGolangBuildpackLatest, additionalOptions: []buildjob.Option{
		buildjob.Since(tester.Release14),
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
