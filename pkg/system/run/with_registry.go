package run

import "github.com/funtimecoding/go-library/pkg/face"

func (r *Run) WithRegistry(registry face.ProcessRegistry) *Run {
	r.registry = registry

	return r
}
