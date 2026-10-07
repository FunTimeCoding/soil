package reference

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func underBases(
	span string,
	bases []string,
	r *resolver.Resolver,
) *Finding {
	var closest *pointer.Resolution

	for _, base := range bases {
		v := r.Resolve(
			"",
			pointer.NewDeclared(),
			pointer.NewSpan(join.Empty(base, stringsConstant.Slash, span)),
		)

		if v.Verdict == lintConstant.VerdictLive ||
			v.Verdict == lintConstant.VerdictTallied {
			return nil
		}

		if closest == nil || closest.Verdict == lintConstant.VerdictDead {
			closest = v
		}
	}

	return NewFinding(span, pointer.Message(closest.Verdict, closest.Hint))
}
