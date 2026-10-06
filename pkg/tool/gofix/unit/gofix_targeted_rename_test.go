package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/unit/module_tester"
	"path/filepath"
	"testing"
)

func TestScopedExportedRenameLeavesNamesake(t *testing.T) {
	for _, full := range []bool{false, true} {
		directory := module_tester.NamesakeExported(t)
		gofix.RunDefault(
			module_tester.Options(t, directory, false, full, "./alfa/..."),
			output.NewResultsWithDirectory(directory),
		)
		assert.String(
			t,
			"package bravo\n\nimport \"testmodule/alfa\"\n\nfunc Relay() string {\n\treturn alfa.SendMessage()\n}\n",
			testutil.ReadFile(t, filepath.Join(directory, "bravo", "relay.go")),
		)
		assert.String(
			t,
			"package charlie\n\nfunc SendMsg() string {\n\treturn \"c\"\n}\n",
			testutil.ReadFile(t, filepath.Join(directory, "charlie", "own.go")),
		)
	}
}
