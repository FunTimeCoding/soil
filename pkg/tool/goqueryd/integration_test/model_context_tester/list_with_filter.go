package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (t *Tester) ListWithFilter(
	collection string,
	sourceType string,
	limit int,
	full bool,
) string {
	arguments := map[string]any{
		constant.Collection: collection,
	}

	if sourceType != "" {
		arguments[constant.SourceType] = sourceType
	}

	if limit > 0 {
		arguments[parameter.Limit] = limit
	}

	if full {
		arguments[constant.Full] = true
	}

	return t.MustCallTool(constant.List, arguments)
}
