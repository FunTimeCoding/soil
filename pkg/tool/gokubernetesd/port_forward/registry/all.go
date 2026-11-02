package registry

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"
	"maps"
	"slices"
)

func (r *Registry) All() []*port_forward.Forward {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return slices.Collect(maps.Values(r.forwards))
}
