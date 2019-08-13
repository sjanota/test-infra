package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestConnectivityCertsController(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("connectivity-certs-controller", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

