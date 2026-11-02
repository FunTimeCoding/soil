package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	result := []face.Worker{d.guard}

	if d.watcher != nil {
		result = append(result, d.rosterGuard, d.watcher)
	}

	for _, c := range d.clocks {
		result = append(result, c)
	}

	return result
}
