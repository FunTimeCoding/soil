package option

import "github.com/funtimecoding/soil/pkg/console/constant"

func New(v ...Option) *Output {
	result := &Output{Format: constant.FormatText, Debug: false}

	for _, o := range v {
		o(result)
	}

	return result
}
