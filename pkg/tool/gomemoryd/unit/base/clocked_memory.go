package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (s *Server) ClockedMemory() int64 {
	s.t.Helper()
	m, e := s.MemoryClient(s.t).PostMemory(
		s.t.Context(),
		client.PostMemoryJSONRequestBody{
			Name:        "clocked",
			Content:     "clocked content",
			Description: "clocked memory",
			Type:        new("feedback"),
		},
	)
	assert.FatalOnError(s.t, e)

	return m.Identifier
}
