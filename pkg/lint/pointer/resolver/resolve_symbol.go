package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"strings"
)

func (r *Resolver) resolveSymbol(candidate string) *pointer.Resolution {
	reference := strings.TrimPrefix(candidate, constant.SchemeGo)
	whole := pointer.Normalize(reference)
	directory := pointer.PackageDirectory(reference)

	if r.Exists(directory) || r.SiblingExists(directory) {
		name := pointer.SymbolName(reference, directory)

		if name != "" && !r.Declares(directory, name) {
			result := pointer.NewResolution(constant.VerdictDead)
			result.Hint = constant.UndeclaredSymbolHint

			return result
		}

		return pointer.NewResolution(constant.VerdictLive)
	}

	if r.Stdlib(directory) ||
		r.Exists(whole) ||
		r.SiblingExists(whole) ||
		r.Stdlib(whole) ||
		r.Ignored(directory) {
		return pointer.NewResolution(constant.VerdictLive)
	}

	return pointer.NewResolution(constant.VerdictDead)
}
