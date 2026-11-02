package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	if d.netboxGuard == nil {
		return []face.Worker{d.worker}
	}

	return []face.Worker{d.netboxGuard, d.worker, d.reconciler}
}
