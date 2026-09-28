package outpost

import "github.com/funtimecoding/soil/pkg/log/logger"

func New(
	targets []*Target,
	l *logger.Logger,
) *Collector {
	return &Collector{targets: targets, logger: l}
}
