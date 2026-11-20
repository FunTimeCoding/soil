package module_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) PackageMoveDocument() *service_tester.Tester {
	o.t.Helper()
	result := service_tester.New(
		o.t,
		service_tester.WithFixture("package-move/src"),
	)
	system_tester.New(o.t).Write(
		result.Directory,
		"doc/note.md",
		"Store lives in `pkg/outer/store/store.go`, built by `go:pkg/outer/store.New`, called from `pkg/caller/run.go`.\n",
	)

	return result
}
