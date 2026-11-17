package base

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func (s *Server) SeedMemory() {
	s.t.Helper()
	s.Tools().MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "alfa",
			constant.Content:     "first",
			constant.Description: "a test",
		},
	)
}
