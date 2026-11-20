package layer_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
)

func (o *Tester) Module() string {
	o.t.Helper()
	root, e := filepath.EvalSymlinks(o.t.TempDir())
	assert.FatalOnError(o.t, e)
	testutil.WriteModFile(o.t, root, "example")
	system_tester.New(o.t).Write(root, "alfa/alfa.go", "package alfa\n")
	system_tester.New(o.t).Write(
		root,
		"alfa/bravo.go",
		"package alfa\n\nfunc Bravo() {}\n",
	)
	system_tester.New(o.t).Write(
		root,
		"charlie/charlie.go",
		"package charlie\n",
	)
	system_tester.New(o.t).Write(
		root,
		"delta/delta.go",
		"package delta\n\nimport \"example/alfa\"\n\nvar _ = alfa.Bravo\n",
	)
	system_tester.New(o.t).Write(root, "doc/echo.md", "echo\n")

	return root
}
