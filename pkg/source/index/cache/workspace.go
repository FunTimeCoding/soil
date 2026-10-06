package cache

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"slices"
)

func (c *Cache) Workspace(root string) *index.Workspace {
	c.lock.Lock()
	cached := c.entries[root]
	c.lock.Unlock()

	if cached != nil && snapshot.Take(cached.trees...).Same(cached.snapshot) {
		return cached.workspace
	}

	trees := module_graph.Trees(root)
	taken := snapshot.Take(trees...)
	var w *index.Workspace

	if cached != nil && slices.Equal(cached.trees, trees) {
		w = index.Next(cached.workspace, cached.snapshot.Diff(taken))
	} else {
		w = index.New(c.directory, root, c.kinds...)
	}

	fresh := &entry{trees: trees, snapshot: taken, workspace: w}
	c.lock.Lock()
	c.entries[root] = fresh
	c.lock.Unlock()

	return fresh.workspace
}
