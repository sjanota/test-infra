package kyma

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
	{path: "acceptance", image: tester.ImageGolangBuildpackLatest},
	{path: "application-connector-tests", image: tester.ImageGolangBuildpackLatest},
	{path: "application-gateway-tests", image: tester.ImageGolangBuildpackLatest},
	{path: "application-operator-tests", image: tester.ImageGolangBuildpackLatest},
	{path: "application-registry-tests", image: tester.ImageGolangBuildpackLatest},
	{path: "asset-store", image: tester.ImageGolangBuildpack1_11},
	{path: "compass-runtime-agent", image: tester.ImageGolangBuildpack1_11},
	{path: "connection-token-handler-tests", image: tester.ImageGolangBuildpackLatest},
	{path: "connector-service-tests", image: tester.ImageGolangBuildpackLatest},
	{path: "console-backend-service", image: tester.ImageGolangBuildpack1_11},
	{path: "end-to-end/backup-restore-test", image: tester.ImageGolangBuildpack1_11},
	{path: "end-to-end/external-solution-integration", image: tester.ImageGolangBuildpack1_11},
	{path: "end-to-end/kubeless-integration", image: tester.ImageGolangBuildpackLatest},
	{path: "end-to-end/upgrade", image: tester.ImageGolangBuildpack1_11, additionalOptions: []buildjob.Option{
		buildjob.RunIfChanged("^tests/end-to-end/upgrade/[^chart]", "tests/end-to-end/upgrade/fix"),
	}},
	{path: "event-bus", image: tester.ImageGolangBuildpack1_11},
	{path: "integration/api-controller", image: tester.ImageGolangBuildpackLatest},
	{path: "integration/apiserver-proxy", image: tester.ImageGolangBuildpackLatest},
	{path: "integration/cluster-users", image: tester.ImageBootstrapLatest},
	{path: "integration/dex", image: tester.ImageGolangBuildpackLatest},
	{path: "integration/event-service", image: tester.ImageGolangBuildpackLatest},
	{path: "integration/gateway", image: tester.ImageGolangBuildpackLatest},
	{path: "integration/logging", image: tester.ImageGolangBuildpackLatest, additionalOptions: []buildjob.Option{
		buildjob.Since(tester.Release14),
	}},
	{path: "integration/monitoring", image: tester.ImageGolangBuildpackLatest, additionalOptions: []buildjob.Option{
		buildjob.Since(tester.Release14),
	}},
	{path: "knative-build", image: tester.ImageGolangBuildpackLatest},
	{path: "knative-serving", image: tester.ImageGolangBuildpackLatest},
	{path: "kubeless", image: tester.ImageGolangBuildpackLatest},
	{path: "logging", image: tester.ImageGolangBuildpackLatest, additionalOptions: []buildjob.Option{
		buildjob.Until(tester.Release13),
	}},
	{path: "monitoring", image: tester.ImageGolangBuildpackLatest, additionalOptions: []buildjob.Option{
		buildjob.Until(tester.Release13),
	}},
	{path: "test-namespace-controller", image: tester.ImageGolangBuildpackLatest},
}

func TestTestJobs(t *testing.T) {
	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			opts := []buildjob.Option{
				buildjob.Test(test.path, test.image),
				buildjob.KymaRepo(),
			}
			opts = append(opts, test.additionalOptions...)
			buildjob.NewSuite(opts...).Run(t)
		})
	}
}
