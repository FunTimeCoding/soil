package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (r *Resolver) resolveRoute(
	bases []string,
	candidate string,
) *Resolution {
	route := Normalize(strings.TrimPrefix(candidate, constant.SchemeRoute))
	route, _, _ = strings.Cut(route, "?")

	if len(bases) == 0 {
		return &Resolution{
			Verdict: constant.VerdictTallied,
			Reason:  constant.ReasonRoute,
		}
	}

	needle := RouteLiteral(route)

	for _, base := range bases {
		if strings.HasPrefix(route, constant.RestRoutePrefix) {
			if paths, found := r.Routes(base); found && MatchRoute(route, paths) {
				return &Resolution{Verdict: constant.VerdictLive}
			}
		}

		if r.Literal(base, needle) {
			return &Resolution{Verdict: constant.VerdictLive}
		}
	}

	for _, base := range r.ImplicitBases {
		if r.Literal(base, needle) {
			return &Resolution{Verdict: constant.VerdictLive}
		}
	}

	return &Resolution{Verdict: constant.VerdictDead}
}
