package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (d *Daemon) Workers() []face.Worker {
	result := []face.Worker{d.prometheusGuard}

	if d.argocdGuard != nil {
		result = append(result, d.argocdGuard)
	}

	if d.usageGuard != nil {
		result = append(result, d.usageGuard)
	}

	return append(
		result,
		worker.NewEvery(
			d.service.Refresh,
			d.logger,
			d.reporter,
			constant.RefreshInterval,
		),
	)
}
