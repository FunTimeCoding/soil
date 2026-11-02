package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"testing"
)

func handMovedFixture(t *testing.T) string {
	t.Helper()
	result := renameFollowFixture(t)
	testutil.WriteFile(
		t,
		result,
		"doc/note.md",
		"See `pkg/alfa/old_name.go`.\n",
	)
	assert.FatalOnError(
		t,
		os.Rename(
			filepath.Join(result, "pkg/alfa/old_name.go"),
			filepath.Join(result, "pkg/alfa/renamed.go"),
		),
	)

	return result
}
