package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func (r *Resolver) resolveRepository(candidate string) *pointer.Resolution {
	normalized := pointer.Normalize(candidate)

	if r.Exists(normalized) || r.Ignored(normalized) {
		return pointer.NewLive(normalized)
	}

	return pointer.NewConcern(constant.DeadPointerKey)
}
