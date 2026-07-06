package lifecycle

import "github.com/funtimecoding/soil/pkg/log/logger"

func New(
	l *logger.Logger,
	v ...Option,
) *Lifecycle {
	result := &Lifecycle{logger: l}

	for _, f := range v {
		f(result)
	}

	return result
}
