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

	if r.SiblingExists(normalized) {
		result := pointer.NewResolution(constant.VerdictLive)
		result.Target = normalized

		return result
	}

	relative, inside := pointer.Relative(path, candidate)

	if inside && r.Exists(relative) {
		result := pointer.NewResolution(constant.VerdictLive)
		result.Target = relative

		return result
	}

	return pointer.NewResolution(constant.VerdictDead)
}
