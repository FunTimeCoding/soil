package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"testing"
)

func updateMatchesBuild(
	t *testing.T,
	root string,
	edit func(),
) *module_graph.Graph {
	t.Helper()
	previous := module_graph.Build(root)
	before := snapshot.Take(module_graph.Trees(root)...)
	edit()
	changed := before.Diff(snapshot.Take(module_graph.Trees(root)...))
	updated := module_graph.Update(previous, root, changed)
	assert.Any(t, module_graph.Build(root), updated)

	return updated
}
