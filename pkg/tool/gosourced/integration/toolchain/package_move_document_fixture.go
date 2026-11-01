package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func packageMoveDocumentFixture(t *testing.T) string {
	t.Helper()
	result := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("package-move/src"),
	)
	testutil.WriteFile(
		t,
		result,
		"doc/note.md",
		"Store lives in `pkg/outer/store/store.go`, built by `go:pkg/outer/store.New`, called from `pkg/caller/run.go`.\n",
	)

	return result
}
