package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) RemoveContext(
	collection string,
	pathPrefix string,
) string {
	return s.Tools().MustCallTool(
		constant.RemoveContext,
		map[string]any{
			constant.Collection: collection,
			constant.PathPrefix: pathPrefix,
		},
	)
}
