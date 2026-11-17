package base

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func (s *Server) RelatedPair() {
	s.t.Helper()
	c := s.Tools()
	c.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "alfa",
			constant.Content:     "alfa content",
			constant.Description: "alfa description",
		},
	)
	c.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "bravo",
			constant.Content:     "bravo content",
			constant.Description: "bravo description",
		},
	)
}
