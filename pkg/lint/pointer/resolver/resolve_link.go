package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func (r *Resolver) resolveLink(
	path string,
	target string,
) *pointer.Resolution {
	relative, inside := pointer.Relative(path, target)

	if inside && r.Exists(relative) {
		return pointer.NewLive(relative)
	}

	return pointer.NewConcern(constant.DeadPointerKey)
}
