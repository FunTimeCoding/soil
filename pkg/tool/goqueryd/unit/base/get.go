package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) Get(path string) string {
	return s.Tools().MustCallTool(
		constant.Get,
		map[string]any{constant.Path: path},
	)
}
