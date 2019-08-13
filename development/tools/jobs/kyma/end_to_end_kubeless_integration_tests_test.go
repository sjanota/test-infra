package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestEndToEndKubelessIntegrationTests(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("end-to-end/kubeless-integration", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

