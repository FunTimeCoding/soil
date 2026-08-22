package detail_error

import "github.com/funtimecoding/soil/pkg/errors/constant"

func (e *Detail) ErrorContext() (string, map[string]any) {
	result := map[string]any{
		constant.Status: e.Status,
		constant.Detail: e.Detail,
	}

	if len(e.Body) > 0 {
		result[constant.Body] = string(e.Body)
	}

	return constant.Upstream, result
}
