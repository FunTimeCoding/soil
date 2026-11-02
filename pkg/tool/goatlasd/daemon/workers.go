package daemon

import "github.com/funtimecoding/soil/pkg/face"

func (d *Daemon) Workers() []face.Worker {
	result := []face.Worker{d.netboxGuard}

	for _, g := range d.guards {
		result = append(result, g)
	}

	return append(result, d.clock)
}
