package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestCmsControllerManager(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("cms-controller-manager", tester.ImageGolangKubebuilderBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

