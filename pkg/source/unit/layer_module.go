package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"path/filepath"
	"testing"
)

func layerModule(t *testing.T) string {
	t.Helper()
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	testutil.WriteFile(t, root, "alfa/alfa.go", "package alfa\n")
	testutil.WriteFile(
		t,
		root,
		"alfa/bravo.go",
		"package alfa\n\nfunc Bravo() {}\n",
	)
	testutil.WriteFile(t, root, "charlie/charlie.go", "package charlie\n")
	testutil.WriteFile(
		t,
		root,
		"delta/delta.go",
		"package delta\n\nimport \"example/alfa\"\n\nvar _ = alfa.Bravo\n",
	)
	testutil.WriteFile(t, root, "doc/echo.md", "echo\n")

	return root
}
