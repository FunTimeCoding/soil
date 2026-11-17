package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) SetTag(
	path string,
	sourceType string,
) string {
	return s.Tools().MustCallTool(
		constant.Tag,
		map[string]any{
			constant.Path:       path,
			constant.SourceType: sourceType,
		},
	)
}
