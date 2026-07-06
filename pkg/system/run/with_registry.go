package run

import "github.com/funtimecoding/soil/pkg/face"

func (r *Run) WithRegistry(registry face.ProcessRegistry) *Run {
	r.registry = registry

	return r
}
