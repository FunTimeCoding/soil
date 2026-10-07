package cache

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"github.com/funtimecoding/soil/pkg/source/types/cache_entry"
	"slices"
)

func (c *Cache) Workspace(root string) *index.Workspace {
	c.lock.Lock()
	cached := c.entries[root]
	c.lock.Unlock()

	if cached != nil && snapshot.Take(cached.Trees...).Same(cached.Snapshot) {
		return cached.Workspace
	}

	trees := module_graph.Trees(root)
	taken := snapshot.Take(trees...)
	var w *index.Workspace

	if cached != nil && slices.Equal(cached.Trees, trees) {
		w = index.Next(cached.Workspace, cached.Snapshot.Diff(taken))
	} else {
		w = index.New(c.directory, root, c.kinds...)
	}

	fresh := cache_entry.New(trees, taken, w)
	c.lock.Lock()
	c.entries[root] = fresh
	c.lock.Unlock()

	return fresh.Workspace
}
