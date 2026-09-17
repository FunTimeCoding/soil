package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (r *Resolver) resolveSibling(
	path string,
	candidate string,
) *Resolution {
	if r.SiblingExists(Normalize(candidate)) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	relative, inside := Relative(path, candidate)

	if inside && r.Exists(relative) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
