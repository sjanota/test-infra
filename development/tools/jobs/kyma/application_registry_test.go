package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestApplicationRegistry(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("application-registry", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

