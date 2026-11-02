package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Stopper() face.Worker {
	return d.stop
}
