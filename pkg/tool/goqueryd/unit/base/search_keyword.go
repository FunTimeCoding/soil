package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) SearchKeyword(query string) string {
	return s.Tools().MustCallTool(
		constant.Search,
		map[string]any{"query": query, constant.Mode: constant.KeywordMode},
	)
}
