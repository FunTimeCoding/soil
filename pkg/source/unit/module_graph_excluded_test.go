package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestExcludedNestedModule(t *testing.T) {
	root := t.TempDir()
	system_tester.New(t).Write(root, "go.mod", "module alfa\n")
	system_tester.New(t).Write(root, "graft/bravo/go.mod", "module bravo\n")
	system_tester.New(t).Write(root, "pkg/alfa/alfa.go", "package alfa\n")
	system_tester.New(t).Write(root, "pkg/alfa/testdata/x.txt", "x\n")
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
