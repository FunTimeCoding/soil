package model_context

import "github.com/funtimecoding/soil/pkg/chromium/snapshot"

func (s *Server) cacheSnapshot(
	tabIdentifier string,
	nodes []*snapshot.Node,
) {
	cache := make(map[string]int64)
	var walk func([]*snapshot.Node)
	walk = func(n []*snapshot.Node) {
		for _, node := range n {
			if node.BackendDOMNodeIdentifier > 0 {
				cache[node.UID] = node.BackendDOMNodeIdentifier
			}

			walk(node.Children)
		}
	}
	walk(nodes)
	s.mutex.Lock()
	s.snapshotCache[tabIdentifier] = cache
	s.mutex.Unlock()
}
