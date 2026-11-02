package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
)

func New(
	o *option.Agent,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	n := runner.New(o.Workspace, o.Stream, l, r)

	return &Daemon{runner: n, service: service.New(n)}
}
