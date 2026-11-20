package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestDefaultRewritesConstantImport(t *testing.T) {
	m := module_tester.New(t)
	directory := m.ConstantAlias()
	r := output.NewResultsWithDirectory(directory)
	m.RunDefault(m.Options(directory, false, false, "./..."), r)
	assert.String(
		t,
		"package charlie\n\nimport bravoConstant \"testmodule/pkg/bravo/constant\"\n\nfunc Lone() string {\n\treturn bravoConstant.Bravo\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "pkg/charlie/lone.go")),
	)
	fix_tester.New(t).Result(
		fix_tester.Applied(r.Entries),
		"pkg/charlie/lone.go",
		"re-aliased constant → bravoConstant (testmodule/pkg/bravo/constant)",
	)

	for _, c := range r.Entries {
		if c.Planned {
			t.Errorf("unexpected suggestion %q on %s", c.Text, c.Path)
		}
	}
}
