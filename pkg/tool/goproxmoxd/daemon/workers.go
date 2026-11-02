package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
)

func (d *Daemon) Workers() []face.Worker {
	var result []face.Worker

	for _, i := range d.inventory.Instances {
		p, e := d.pool.Client(i.Name)
		errors.PanicOnError(e)
		result = append(result, p.Guard())
	}

	if d.netboxGuard == nil {
		return append(result, d.worker)
	}

	return append(result, d.netboxGuard, d.worker, d.reconciler)
}
