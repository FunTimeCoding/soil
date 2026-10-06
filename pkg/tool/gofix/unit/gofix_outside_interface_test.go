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

func TestScopedRunKeepsMethodOfOutsideInterface(t *testing.T) {
	for _, full := range []bool{false, true} {
		directory := module_tester.OutsideInterface(t)
		gofix.RunDefault(
			module_tester.Options(t, directory, false, full, "./alfa/..."),
			output.NewResultsWithDirectory(directory),
		)
		assert.String(
			t,
			"package alfa\n\ntype Sender struct{}\n\nfunc (Sender) SendMsg() string {\n\treturn \"x\"\n}\n",
			testutil.ReadFile(t, filepath.Join(directory, "alfa", "sender.go")),
		)
		assert.String(
			t,
			"package bravo\n\nimport \"testmodule/alfa\"\n\ntype messenger interface {\n\tSendMsg() string\n}\n\nvar Chosen messenger = alfa.Sender{}\n",
			testutil.ReadFile(
				t,
				filepath.Join(directory, "bravo", "messenger.go"),
			),
		)
	}
}
