package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"testing"
)

func TestDefaultFormatsAfterDealiasing(t *testing.T) {
	m := module_tester.New(t)
	directory := m.AliasedCall()
	m.RunDefault(
		m.Options(directory, false, false, "./..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nimport \"testmodule/pkg/helper\"\n\nfunc Use() string {\n\treturn helper.Join(\"first-literal-value\", \"second-literal-value\", \"third\")\n}\n",
		fix_tester.New(t).GoFiles(directory)["use.go"],
	)
}
