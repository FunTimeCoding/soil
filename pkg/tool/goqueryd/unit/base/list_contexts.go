package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) ListContexts() string {
	return s.Tools().MustCallTool(constant.ListContexts, map[string]any{})
}
