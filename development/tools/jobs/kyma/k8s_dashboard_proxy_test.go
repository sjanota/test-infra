package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestK8sDashboardProxy(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("k8s-dashboard-proxy", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

