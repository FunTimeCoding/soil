package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	var result []face.Worker

	for _, g := range d.guards {
		result = append(result, g)
	}

	return append(result, d.reaper, d.runner)
}
