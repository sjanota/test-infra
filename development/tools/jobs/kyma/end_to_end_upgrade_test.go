package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestEndToEndUpgrade(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("end-to-end/upgrade", tester.ImageGolangBuildpack1_11),
		buildjob.KymaRepo(),
		buildjob.RunIfChanged("^tests/end-to-end/upgrade/[^chart]", "tests/end-to-end/upgrade/fix"),
	).Run(t)
}

