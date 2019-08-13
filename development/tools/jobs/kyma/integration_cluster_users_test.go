package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestIntegrationClusterUsers(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("integration/cluster-users", tester.ImageBootstrapLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

