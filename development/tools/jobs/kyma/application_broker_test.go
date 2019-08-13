package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestApplicationBrokerServer(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("application-broker", tester.ImageGolangBuildpack1_11),
		buildjob.KymaRepo(),
	).Run(t)
}
