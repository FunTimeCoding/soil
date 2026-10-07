package outpost

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
)

func New(
	targets []*target.Target,
	l *logger.Logger,
) *Collector {
	return &Collector{targets: targets, logger: l}
}
