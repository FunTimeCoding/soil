package registry

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"

func New() *Registry {
	return &Registry{forwards: map[string]*port_forward.Forward{}}
}
