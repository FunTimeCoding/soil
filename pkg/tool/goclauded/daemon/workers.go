package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	result := []face.Worker{d.memoryGuard, d.queryGuard}

	if d.utilizationGuard != nil {
		result = append(result, d.utilizationGuard)
	}

	result = append(result, d.clocks...)

	return append(result, d.gauge, d.watcher)
}
