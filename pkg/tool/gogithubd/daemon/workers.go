package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	return []face.Worker{d.clock}
}
