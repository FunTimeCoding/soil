package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"maps"
	"path/filepath"
	"slices"
	"testing"
)

func TestBuildStopsAtNestedModule(t *testing.T) {
	root := t.TempDir()
	testutil.WriteModFile(t, root, "example.org/alfa")
	testutil.WriteFile(t, root, "bravo/bravo.go", "package bravo\n")
	testutil.WriteFile(t, root, "graft/charlie/charlie.go", "package charlie\n")
	testutil.WriteModFile(
		t,
		filepath.Join(root, "graft", "charlie"),
		"example.org/charlie",
	)
	assert.Strings(
		t,
		[]string{"example.org/alfa/bravo"},
		slices.Sorted(maps.Keys(module_graph.Build(root).Nodes)),
	)
}

func TestBuildForestReadsTreeNotDisk(t *testing.T) {
	root := t.TempDir()
	testutil.WriteModFile(t, root, "example.org/alfa")
	testutil.WriteFile(t, root, "bravo/bravo.go", "package bravo\n")
	testutil.WriteFile(t, root, "charlie/charlie.go", "package charlie\n")
	f := warm_forest.New()
	_, e := f.Tree(root).Read("bravo/bravo.go")
	assert.FatalOnError(t, e)
	testutil.WriteFile(
		t,
		root,
		"bravo/bravo.go",
		"package bravo\n\nimport _ \"example.org/alfa/charlie\"\n",
	)
	assert.Count(
		t,
		0,
		module_graph.BuildForest(f, root).Nodes["example.org/alfa/bravo"].Imports,
	)
}
