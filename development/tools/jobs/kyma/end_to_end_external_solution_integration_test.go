package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestEndToEndExternalSolutionIntegration(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("end-to-end/external-solution-integration", tester.ImageGolangBuildpack1_11),
		buildjob.KymaRepo(),
	).Run(t)
}

