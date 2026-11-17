package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) Push(
	collection string,
	path string,
	body string,
) string {
	return s.Tools().MustCallTool(
		constant.Push,
		map[string]any{
			constant.Collection: collection,
			constant.Path:       path,
			constant.Body:       body,
		},
	)
}
