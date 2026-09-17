package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (r *Resolver) resolveLink(
	path string,
	target string,
) *Resolution {
	relative, inside := Relative(path, target)

	if inside && r.Exists(relative) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
