package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestAcceptance(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("acceptance", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
	).Run(t)
}
