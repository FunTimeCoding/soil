package scheduler

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/robfig/cron/v3"
)

func New(
	schedule string,
	task func(),
	l *logger.Logger,
	r face.Reporter,
) *Scheduler {
	return &Scheduler{
		schedule: schedule,
		task:     task,
		logger:   l,
		reporter: r,
		cron:     cron.New(),
	}
}
