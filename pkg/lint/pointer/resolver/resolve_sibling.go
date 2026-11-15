package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func (r *Resolver) resolveSibling(
	path string,
	candidate string,
) *pointer.Resolution {
	normalized := pointer.Normalize(candidate)

	if r.SiblingExists(normalized) || r.Ignored(normalized) {
		return pointer.NewLive(normalized)
	}

	relative, inside := pointer.Relative(path, candidate)

	if inside && r.Exists(relative) {
		return pointer.NewLive(relative)
	}

	return pointer.NewConcern(constant.DeadPointerKey)
}
