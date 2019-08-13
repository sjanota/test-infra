package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestXipPatch(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("xip-patch", tester.ImageBootstrapLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

