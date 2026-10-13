package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"path/filepath"
	"testing"
)

func TestExcludedNestedModule(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "go.mod", "module alfa\n")
	testutil.WriteFile(t, root, "graft/bravo/go.mod", "module bravo\n")
	testutil.WriteFile(t, root, "pkg/alfa/alfa.go", "package alfa\n")
	testutil.WriteFile(t, root, "pkg/alfa/testdata/x.txt", "x\n")
	assert.Boolean(t, false, module_graph.Excluded(root, root))
	assert.True(
		t,
		module_graph.Excluded(root, filepath.Join(root, "graft", "bravo")),
	)
	assert.Boolean(
		t,
		false,
		module_graph.Excluded(root, filepath.Join(root, "pkg", "alfa")),
	)
	assert.True(
		t,
		module_graph.Excluded(
			root,
			filepath.Join(root, "pkg", "alfa", "testdata"),
		),
	)
}
