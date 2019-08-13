package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestServiceBindingUsageController(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("service-binding-usage-controller", tester.ImageGolangBuildpack1_11),
		buildjob.KymaRepo(),
	).Run(t)
}

