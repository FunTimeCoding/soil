package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
)

func (o *Tester) HandMoved() *service_tester.Tester {
	o.t.Helper()
	result := service_tester.New(
		o.t,
		service_tester.WithFixture("rename-follow/src"),
	)
	system_tester.New(o.t).Write(
		result.Directory,
		"doc/note.md",
		"See `pkg/alfa/old_name.go`.\n",
	)
	assert.FatalOnError(
		o.t,
		os.Rename(
			filepath.Join(result.Directory, "pkg/alfa/old_name.go"),
			filepath.Join(result.Directory, "pkg/alfa/renamed.go"),
		),
	)

	return result
}
