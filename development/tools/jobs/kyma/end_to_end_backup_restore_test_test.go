package kyma_test

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
)

func TestEndToEndBackupRestoreTest(t *testing.T) {
	buildjob.NewSuite(
		buildjob.Test("end-to-end/backup-restore-test", tester.ImageGolangBuildpack1_11),
		buildjob.KymaRepo(),
	).Run(t)
}

