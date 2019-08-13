package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestKnativeFunctionController(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("knative-function-controller", tester.ImageGolangKubebuilderBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

