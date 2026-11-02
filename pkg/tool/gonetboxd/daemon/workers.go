package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	if d.clock == nil {
		return nil
	}

	return []face.Worker{d.clock}
}
