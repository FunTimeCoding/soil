package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"maps"
	"slices"
)

func (d *Daemon) Workers() []face.Worker {
	var result []face.Worker

	for _, name := range slices.Sorted(maps.Keys(d.guards)) {
		result = append(result, d.guards[name])
	}

	return append(result, d.forwards)
}
