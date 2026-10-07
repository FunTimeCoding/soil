package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func (r *Resolver) resolveRepository(candidate string) *pointer.Resolution {
	normalized := pointer.Normalize(candidate)

	if r.Exists(normalized) || r.Ignored(normalized) {
		result := pointer.NewResolution(constant.VerdictLive)
		result.Target = normalized

		return result
	}

	return pointer.NewResolution(constant.VerdictDead)
}
