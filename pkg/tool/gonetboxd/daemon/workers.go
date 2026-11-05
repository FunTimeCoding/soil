package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	if d.clock == nil {
		return []face.Worker{d.guard}
	}

	return []face.Worker{d.guard, d.clock}
}
