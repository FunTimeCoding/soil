package check_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
)

func (o *Tester) StackModule() string {
	o.t.Helper()
	root, e := filepath.EvalSymlinks(o.t.TempDir())
	assert.FatalOnError(o.t, e)
	testutil.WriteModFile(o.t, root, "example")
	system_tester.New(o.t).Write(
		root,
		"pkg/tool/charlie/charlie.go",
		"package charlie\n",
	)
	system_tester.New(o.t).Write(
		root,
		"cmd/charlie/main.go",
		"package main\n\nfunc main() {}\n",
	)
	system_tester.New(o.t).Write(
		root,
		"pkg/spacing/alfa.go",
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
	)
	system_tester.New(o.t).Write(
		root,
		"pkg/format/delta.go",
		"package format\n\nimport \"fmt\"\n\nfunc Delta() {\n\tfmt.Println(\"echo foxtrot golf hotel india\", \"juliett kilo lima mike november oscar\", \"x\")\n}\n",
	)

	return root
}
