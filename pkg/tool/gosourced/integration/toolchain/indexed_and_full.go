package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func indexedAndFull(
	t *testing.T,
) (string, *service_tester.Services, *service_tester.Services) {
	t.Helper()
	directory := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("indexed/src"),
	)
	full := service_tester.NewService(t)
	full.Workspace.UseFullLoad()

	return directory, service_tester.NewService(t), full
}
