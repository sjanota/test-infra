package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestDexStaticUserConfigurer(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Component("dex-static-user-configurer", tester.ImageBootstrapLatest),
		buildjob.KymaRepo(),
	).Run(t)
}

