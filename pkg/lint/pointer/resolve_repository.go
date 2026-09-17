package pointer

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (r *Resolver) resolveRepository(candidate string) *Resolution {
	normalized := Normalize(candidate)

	if r.Exists(normalized) || r.Ignored(normalized) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
