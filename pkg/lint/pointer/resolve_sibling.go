package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (r *Resolver) resolveSibling(
	path string,
	candidate string,
) *Resolution {
	normalized := Normalize(candidate)

	if r.SiblingExists(normalized) {
		return &Resolution{Verdict: constant.VerdictLive, Target: normalized}
	}

	relative, inside := Relative(path, candidate)

	if inside && r.Exists(relative) {
		return &Resolution{Verdict: constant.VerdictLive, Target: relative}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
