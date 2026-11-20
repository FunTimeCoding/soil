package module_graph_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
)

func (o *Tester) Updated(
	root string,
	edit func(),
) *module_graph.Graph {
	o.t.Helper()
	previous := module_graph.Build(root)
	before := snapshot.Take(module_graph.Trees(root)...)
	edit()
	changed := before.Diff(snapshot.Take(module_graph.Trees(root)...))
	updated := module_graph.Update(previous, root, changed)
	assert.Any(o.t, module_graph.Build(root), updated)

	return updated
}
