package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestApiserverProxyIntegrationTests(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("integration/apiserver-proxy", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}
