package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestApiServer(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("apiserver-proxy", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}
