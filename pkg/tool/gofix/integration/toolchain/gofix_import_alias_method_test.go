package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestImportAliasFixIgnoresMethodNamesake(t *testing.T) {
	directory := module_tester.New(t).MethodNamesake()
	gofix.RunImportAliasFix(
		[]string{"./..."},
		directory,
		false,
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package example\n\nimport \"testmodule/pkg/helper\"\n\ntype Holder struct{}\n\nfunc (o *Holder) helper() string {\n\treturn helper.Value()\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "holder.go")),
	)
	fix_tester.New(t).Builds(directory)
}
