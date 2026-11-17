package module_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func PackageMoveDocument(t *testing.T) *service_tester.Tester {
	t.Helper()
	result := service_tester.New(
		t,
		service_tester.WithFixture("package-move/src"),
	)
	system_tester.Write(
		t,
		result.Directory,
		"doc/note.md",
		"Store lives in `pkg/outer/store/store.go`, built by `go:pkg/outer/store.New`, called from `pkg/caller/run.go`.\n",
	)

	return result
}
