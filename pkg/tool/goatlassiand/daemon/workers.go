package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (d *Daemon) Workers() []face.Worker {
	return []face.Worker{
		d.jiraGuard,
		d.confluenceGuard,
		worker.NewEvery(
			d.service.RefreshPlate,
			d.logger,
			d.reporter,
			constant.PollInterval,
		),
	}
}
