package check_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func CacheModule(t *testing.T) string {
	t.Helper()
	root, e := filepath.EvalSymlinks(t.TempDir())
	assert.FatalOnError(t, e)
	testutil.WriteModFile(t, root, "example")
	system_tester.Write(
		t,
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Alfa() int {\n\treturn 1\n}\n",
	)
	system_tester.Write(
		t,
		root,
		"pkg/bravo/bravo.go",
		"package bravo\n\nimport \"example/pkg/alfa\"\n\nfunc Bravo() int {\n\tprintln(\"delta echo\")\n\n\treturn alfa.Alfa()\n}\n",
	)
	system_tester.Write(
		t,
		root,
		"pkg/response/types.go",
		"package response\n\ntype LinkResponse struct {\n\tName string\n}\n",
	)

	return root
}
