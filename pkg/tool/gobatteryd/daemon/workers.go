package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	var result []face.Worker

	for _, t := range d.targets {
		result = append(result, t.Guard)
	}

	return append(result, d.clock)
}
