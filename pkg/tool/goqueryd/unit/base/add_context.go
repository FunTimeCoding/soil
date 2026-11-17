package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) AddContext(
	collection string,
	pathPrefix string,
	description string,
) string {
	return s.Tools().MustCallTool(
		constant.AddContext,
		map[string]any{
			constant.Collection:  collection,
			constant.PathPrefix:  pathPrefix,
			constant.Description: description,
		},
	)
}
