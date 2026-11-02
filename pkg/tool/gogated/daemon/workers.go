package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (d *Daemon) Workers() []face.Worker {
	var result []face.Worker

	if d.directoryGuard != nil {
		result = append(result, d.directoryGuard)
	}

	return append(
		result,
		worker.NewEvery(
			d.service.Clean,
			d.logger,
			d.reporter,
			constant.CleanInterval,
		),
	)
}
