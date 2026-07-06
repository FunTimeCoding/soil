package recovery

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(
	l *logger.Logger,
	r face.Reporter,
) *Recovery {
	return &Recovery{logger: l, reporter: r}
}
