package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestDefaultRewritesConstantImport(t *testing.T) {
	directory := module_tester.ConstantAlias(t)
	r := output.NewResultsWithDirectory(directory)
	module_tester.RunDefault(
		t,
		module_tester.Options(t, directory, false, false, "./..."),
		r,
	)
	assert.String(
		t,
		"package charlie\n\nimport bravoConstant \"testmodule/pkg/bravo/constant\"\n\nfunc Lone() string {\n\treturn bravoConstant.Bravo\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "pkg/charlie/lone.go")),
	)
	assertResult(
		t,
		filterApplied(r.Entries),
		"pkg/charlie/lone.go",
		"re-aliased constant → bravoConstant (testmodule/pkg/bravo/constant)",
	)

	for _, c := range r.Entries {
		if c.Planned {
			t.Errorf("unexpected suggestion %q on %s", c.Text, c.Path)
		}
	}
}
