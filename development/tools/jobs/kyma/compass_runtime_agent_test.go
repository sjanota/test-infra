package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestCompassRuntimeAgent(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("compass-runtime-agent", tester.ImageGolangBuildpack1_11),
		buildjob.KymaRepo(),
	).Run(t)
}

