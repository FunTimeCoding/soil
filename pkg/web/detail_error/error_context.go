package detail_error

import errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"

func (e *Detail) ErrorContext() (string, map[string]any) {
	result := map[string]any{
		errorsConstant.Status: e.Status,
		errorsConstant.Detail: e.Detail,
	}

	if len(e.Body) > 0 {
		result[errorsConstant.Body] = string(e.Body)
	}

	return errorsConstant.Upstream, result
}
