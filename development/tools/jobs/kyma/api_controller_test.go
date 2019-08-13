package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestApiController(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("api-controller", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}