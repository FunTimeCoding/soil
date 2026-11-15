package layer_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func Module(t *testing.T) string {
	t.Helper()
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	system_tester.Write(t, root, "alfa/alfa.go", "package alfa\n")
	system_tester.Write(
		t,
		root,
		"alfa/bravo.go",
		"package alfa\n\nfunc Bravo() {}\n",
	)
	system_tester.Write(t, root, "charlie/charlie.go", "package charlie\n")
	system_tester.Write(
		t,
		root,
		"delta/delta.go",
		"package delta\n\nimport \"example/alfa\"\n\nvar _ = alfa.Bravo\n",
	)
	system_tester.Write(t, root, "doc/echo.md", "echo\n")

	return root
}
