package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/unit/module_tester"
	"testing"
)

func TestDefaultFormatsAfterDealiasing(t *testing.T) {
	directory := module_tester.AliasedCall(t)
	gofix.RunDefault(
		module_tester.Options(t, directory, false, false, "./..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nimport \"testmodule/pkg/helper\"\n\nfunc Use() string {\n\treturn helper.Join(\"first-literal-value\", \"second-literal-value\", \"third\")\n}\n",
		readGoFiles(t, directory)["use.go"],
	)
}
