package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"io"
)

func New(
	workspace string,
	stream func(string, io.Writer, ...string) face.Stream,
	l *logger.Logger,
	r face.Reporter,
) *Runner {
	return &Runner{
		state:     constant.RunnerIdle,
		workspace: workspace,
		stream:    stream,
		reporter:  r,
		recovery:  recovery.New(l, r),
		finished:  make(chan struct{}),
	}
}
