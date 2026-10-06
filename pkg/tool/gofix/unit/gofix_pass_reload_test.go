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

func TestPassAfterRenameSeesRenamedSource(t *testing.T) {
	directory := module_tester.RenamedParameter(t)
	gofix.RunDefault(
		module_tester.Options(t, directory, false, false, "./..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nfunc Join(m string) string {\n\treturn m\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "join.go")),
	)
}
