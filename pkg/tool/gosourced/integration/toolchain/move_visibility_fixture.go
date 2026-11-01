package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func moveVisibilityFixture(t *testing.T) string {
	t.Helper()

	return testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-visibility/src"),
	)
}
