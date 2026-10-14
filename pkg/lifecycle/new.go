package lifecycle

import (
	"github.com/funtimecoding/soil/pkg/lifecycle/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(l *logger.Logger) *Lifecycle {
	return &Lifecycle{logger: l, grace: constant.TerminationGrace}
}
