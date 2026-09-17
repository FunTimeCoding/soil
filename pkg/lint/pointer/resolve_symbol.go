package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (r *Resolver) resolveSymbol(candidate string) *Resolution {
	reference := strings.TrimPrefix(candidate, constant.SchemeGo)
	whole := Normalize(reference)
	directory := PackageDirectory(reference)

	if r.Exists(directory) ||
		r.SiblingExists(directory) ||
		r.Stdlib(directory) ||
		r.Exists(whole) ||
		r.SiblingExists(whole) ||
		r.Stdlib(whole) ||
		r.Ignored(directory) {
		return &Resolution{Verdict: constant.VerdictLive}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
