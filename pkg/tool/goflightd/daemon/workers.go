package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	var result []face.Worker

	for _, c := range d.clocks {
		result = append(result, c)
	}

	return result
}
