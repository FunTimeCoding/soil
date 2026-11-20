package check_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
)

func (o *Tester) CacheModule() string {
	o.t.Helper()
	root, e := filepath.EvalSymlinks(o.t.TempDir())
	assert.FatalOnError(o.t, e)
	testutil.WriteModFile(o.t, root, "example")
	system_tester.New(o.t).Write(
		root,
		"pkg/alfa/alfa.go",
		"package alfa\n\nfunc Alfa() int {\n\treturn 1\n}\n",
	)
	system_tester.New(o.t).Write(
		root,
		"pkg/bravo/bravo.go",
		"package bravo\n\nimport \"example/pkg/alfa\"\n\nfunc Bravo() int {\n\tprintln(\"delta echo\")\n\n\treturn alfa.Alfa()\n}\n",
	)
	system_tester.New(o.t).Write(
		root,
		"pkg/response/types.go",
		"package response\n\ntype LinkResponse struct {\n\tName string\n}\n",
	)

	return root
}
