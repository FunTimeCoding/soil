package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Server) visibleMemories(memories []store.Memory) []store.Memory {
	result := make([]store.Memory, 0, len(memories))

	for _, m := range memories {
		if s.skipHidden(m.Tags) {
			continue
		}

		result = append(result, m)
	}

	return result
}
