package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestHelmBrokerDeprecated(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("helm-broker", tester.ImageGolangKubebuilderBuildpackLatest),
		buildjob.JobFileSuffix("deprecated"),
		buildjob.KymaRepo(),
		buildjob.Until(tester.Release14),
	).Run(t)
}

