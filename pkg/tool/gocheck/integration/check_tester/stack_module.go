package check_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func StackModule(t *testing.T) string {
	t.Helper()
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	system_tester.Write(
		t,
		root,
		"pkg/tool/charlie/charlie.go",
		"package charlie\n",
	)
	system_tester.Write(
		t,
		root,
		"cmd/charlie/main.go",
		"package main\n\nfunc main() {}\n",
	)
	system_tester.Write(
		t,
		root,
		"pkg/spacing/alfa.go",
		"package spacing\n\nfunc Alfa() int {\n\tbravo := 1\n\treturn bravo\n}\n",
	)
	system_tester.Write(
		t,
		root,
		"pkg/format/delta.go",
		"package format\n\nimport \"fmt\"\n\nfunc Delta() {\n\tfmt.Println(\"echo foxtrot golf hotel india\", \"juliett kilo lima mike november oscar\", \"x\")\n}\n",
	)

	return root
}
