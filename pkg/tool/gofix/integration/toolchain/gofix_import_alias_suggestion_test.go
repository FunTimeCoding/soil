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
	gofix.RunImportAliasFix([]string{"./..."}, directory, false, r)
	t.Run(
		"StandardLibraryKeepsThePlainName",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/delta/clock.go",
				"suggest alias library → testmoduleTime (testmodule/pkg/time; time keeps the name: stdlib)",
			)
		},
	)
	t.Run(
		"OwnKeepsThePlainName",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/echo/service/service.go",
				"suggest de-alias echoFace → face (own: testmodule/pkg/echo/face)",
			)
			assertPlanned(
				t,
				r.Entries,
				"pkg/echo/service/service.go",
				"suggest alias face → foxtrot (testmodule/pkg/foxtrot/face; testmodule/pkg/echo/face keeps the name: own)",
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
}

func TestOtherCollisionsStayUntouchedAndSilentByDefault(t *testing.T) {
	directory := module_tester.ImportAliasSuggestion(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunConstantAliasFix([]string{"./..."}, directory, false, r)
	assert.String(
		t,
		"package service\n\nimport (\n\techoFace \"testmodule/pkg/echo/face\"\n\t\"testmodule/pkg/foxtrot/face\"\n)\n\nfunc Service(echoFace.Client, face.Source) {}\n",
		testutil.ReadFile(
			t,
			filepath.Join(directory, "pkg/echo/service/service.go"),
		),
	)

	for _, c := range r.Entries {
		if c.Path == "pkg/echo/service/service.go" ||
			c.Path == "pkg/hotel/builtin.go" {
			t.Errorf("unexpected concern %q on %s", c.Text, c.Path)
		}
	}
}
