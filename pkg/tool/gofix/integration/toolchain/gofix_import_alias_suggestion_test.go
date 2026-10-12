package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestImportAliasSuggestion(t *testing.T) {
	directory := module_tester.ImportAliasSuggestion(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunRealiasFixWithDirectory([]string{"./..."}, directory, false, r)
	t.Run(
		"StandardLibraryKeepsThePlainName",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/delta/clock.go",
				"suggest alias library → testmodule (testmodule/pkg/time; time keeps the name: stdlib)",
			)
		},
	)
	t.Run(
		"KeeperRefusesABuiltinName",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/hotel/builtin.go",
				"cannot de-alias own → string (builtin collision)",
			)
		},
	)
	t.Run(
		"OtherCollisionsOnlySuggestEvenWhenRealiasing",
		func(t *testing.T) {
			assert.String(
				t,
				"package delta\n\nimport (\n\tlibrary \"testmodule/pkg/time\"\n\t\"time\"\n)\n\nfunc Clock() string {\n\treturn library.Stamp() + time.UTC.String()\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/delta/clock.go"),
				),
			)
		},
	)
}
