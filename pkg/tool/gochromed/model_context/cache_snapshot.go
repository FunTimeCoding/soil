package model_context

import "github.com/funtimecoding/soil/pkg/chromium/snapshot"

func (s *Server) cacheSnapshot(
	tabID string,
	nodes []*snapshot.Node,
) {
	cache := make(map[string]int64)
	var walk func([]*snapshot.Node)
	walk = func(n []*snapshot.Node) {
		for _, node := range n {
			if node.BackendDOMNodeID > 0 {
				cache[node.UID] = node.BackendDOMNodeID
			}

			walk(node.Children)
		}
	}
	walk(nodes)
	s.mutex.Lock()
	s.snapshotCache[tabID] = cache
	s.mutex.Unlock()
}
