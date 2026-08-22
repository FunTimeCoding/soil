package job

import "github.com/funtimecoding/soil/pkg/errors/constant"

func (e *JobError) ErrorContext() (string, map[string]any) {
	result := map[string]any{
		constant.Identifier: e.Identifier,
		constant.Kind:       e.Kind,
	}

	for key, value := range e.Detail {
		result[key] = value
	}

	return constant.Job, result
}
