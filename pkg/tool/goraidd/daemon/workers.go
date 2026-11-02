package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/time/time_of_day"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/worker"
	"time"
)

func (d *Daemon) Workers() []face.Worker {
	return []face.Worker{
		d.parserGuard,
		d.watcher,
		worker.NewDaily(
			d.service.Cleanup,
			d.logger,
			d.reporter,
			time_of_day.New(
				constant.CleanupHour,
				constant.CleanupMinute,
				time.Local,
			),
		),
	}
}
