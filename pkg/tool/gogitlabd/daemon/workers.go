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

	result = append(result, d.worker)

	if d.rosterGuard != nil {
		result = append(result, d.rosterGuard)
	}

	for _, c := range d.clocks {
		result = append(result, c)
	}

	return result
}
