package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestApplicationConnectivityCertsSetupJob(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("application-connectivity-certs-setup-job", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

