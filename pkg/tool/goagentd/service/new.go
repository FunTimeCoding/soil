package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
)

func New(
	o *option.Agent,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	return &Service{runner: runner.New(o.Workspace, o.Stream, l, r)}
}
