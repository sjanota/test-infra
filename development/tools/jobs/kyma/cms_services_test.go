package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestCmsServices(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("cms-services", tester.ImageGolangBuildpack1_12),
		buildjob.KymaRepo(),
	).Run(t)
}

