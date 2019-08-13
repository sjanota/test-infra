package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestIstioKymaPatch(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("istio-kyma-patch", tester.ImageBootstrapLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

