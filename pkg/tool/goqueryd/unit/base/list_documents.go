package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) ListDocuments(collection string) string {
	return s.Tools().MustCallTool(
		constant.List,
		map[string]any{constant.Collection: collection},
	)
}
