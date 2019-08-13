package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestMonitoring(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("integration/monitoring", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
		buildjob.Since(tester.Release14),
	).Run(t)
}


func TestMonitoringOld(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("monitoring", tester.ImageGolangBuildpackLatest),
		buildjob.KymaRepo(),
		buildjob.Until(tester.Release13),
	).Run(t)
}
