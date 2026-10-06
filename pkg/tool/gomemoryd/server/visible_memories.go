package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Server) visibleMemories(memories []record.Memory) []record.Memory {
	result := make([]record.Memory, 0, len(memories))

	for _, m := range memories {
		if s.skipHidden(m.Tags) {
			continue
		}

		result = append(result, m)
	}

	return result
}
