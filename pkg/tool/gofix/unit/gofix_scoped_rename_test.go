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

func TestScopedExportedRenameReachesCallerTyped(t *testing.T) {
	directory := module_tester.ExportedCaller(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunDefault(
		module_tester.Options(t, directory, false, false, "./alfa/..."),
		r,
	)
	assertResult(
		t,
		filterApplied(r.Entries),
		"alfa/send.go",
		"renamed SendMsg → SendMessage (2 references)",
	)
	assert.String(
		t,
		"package bravo\n\nimport \"testmodule/alfa\"\n\nfunc Relay() string {\n\treturn alfa.SendMessage()\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "bravo", "relay.go")),
	)
}

func TestScopedUnexportedRenameReachesOwnTest(t *testing.T) {
	directory := module_tester.UnexportedWithTests(t)
	gofix.RunDefault(
		module_tester.Options(t, directory, false, false, "./alfa/..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package alfa\n\nfunc sendMessage() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "send.go")),
	)
	assert.String(
		t,
		"package alfa\n\nimport \"testing\"\n\nfunc TestSend(t *testing.T) {\n\t_ = sendMessage()\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "send_test.go")),
	)
}
