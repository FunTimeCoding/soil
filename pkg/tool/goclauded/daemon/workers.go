package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	result := []face.Worker{d.memoryGuard, d.queryGuard}
	result = append(result, d.clocks...)

	return append(result, d.gauge, d.watcher)
}
