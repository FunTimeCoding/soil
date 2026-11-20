package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	return []face.Worker{
		d.nationalGuard,
		d.trackerGuard,
		d.extendedGuard,
		d.clock,
	}
}
