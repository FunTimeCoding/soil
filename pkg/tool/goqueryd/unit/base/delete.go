package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) Delete(
	collection string,
	path string,
) string {
	return s.Tools().MustCallTool(
		constant.Delete,
		map[string]any{
			constant.Collection: collection,
			constant.Path:       path,
		},
	)
}
