package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) GetTag(path string) string {
	return s.Tools().MustCallTool(
		constant.Tag,
		map[string]any{constant.Path: path},
	)
}
