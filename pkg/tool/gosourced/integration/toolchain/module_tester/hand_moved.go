package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func HandMoved(t *testing.T) *service_tester.Tester {
	t.Helper()
	result := service_tester.New(
		t,
		service_tester.WithFixture("rename-follow/src"),
	)
	system_tester.Write(
		t,
		result.Directory,
		"doc/note.md",
		"See `pkg/alfa/old_name.go`.\n",
	)
	assert.FatalOnError(
		t,
		os.Rename(
			filepath.Join(result.Directory, "pkg/alfa/old_name.go"),
			filepath.Join(result.Directory, "pkg/alfa/renamed.go"),
		),
	)

	return result
}
