package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (s *Server) ListWithFilter(
	collection string,
	sourceType string,
	limit int,
	full bool,
) string {
	arguments := map[string]any{constant.Collection: collection}

	if sourceType != "" {
		arguments[constant.SourceType] = sourceType
	}

	if limit > 0 {
		arguments[generativeConstant.ParameterLimit] = limit
	}

	if full {
		arguments[constant.Full] = true
	}

	return s.Tools().MustCallTool(constant.List, arguments)
}
