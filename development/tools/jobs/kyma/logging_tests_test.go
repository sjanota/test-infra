package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestLoggingTests(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("integration/logging", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
		buildjob.Since(tester.Release14),
	).Run(t)
}

func TestLoggingTestsOld(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("logging", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
		buildjob.Until(tester.Release13),
	).Run(t)
}

