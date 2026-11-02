package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	if d.inventoryGuard == nil {
		return []face.Worker{d.reaper, d.runner}
	}

	return []face.Worker{d.inventoryGuard, d.reaper, d.runner}
}
