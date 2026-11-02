package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func renameFunctionFixture(t *testing.T) string {
	t.Helper()

	return testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-function/src"),
	)
}
