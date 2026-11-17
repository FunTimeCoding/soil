package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Server) PushWithMetadata(
	collection string,
	path string,
	body string,
	metadata map[string]string,
) string {
	arguments := map[string]any{
		constant.Collection: collection,
		constant.Path:       path,
		constant.Body:       body,
	}

	if len(metadata) > 0 {
		arguments[constant.Metadata] = metadata
	}

	return s.Tools().MustCallTool(constant.Push, arguments)
}
