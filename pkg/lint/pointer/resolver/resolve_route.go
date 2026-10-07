package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"strings"
)

func (r *Resolver) resolveRoute(
	bases []string,
	candidate string,
) *pointer.Resolution {
	route := pointer.Normalize(
		strings.TrimPrefix(candidate, constant.SchemeRoute),
	)
	route, _, _ = strings.Cut(route, "?")

	if len(bases) == 0 {
		result := pointer.NewResolution(constant.VerdictTallied)
		result.Reason = constant.ReasonRoute

		return result
	}

	needle := pointer.RouteLiteral(route)

	for _, base := range bases {
		if strings.HasPrefix(route, constant.RestRoutePrefix) {
			if paths, found := r.Routes(base); found && pointer.MatchRoute(
				route,
				paths,
			) {
				return pointer.NewResolution(constant.VerdictLive)
			}
		}

		if r.Literal(base, needle) {
			return pointer.NewResolution(constant.VerdictLive)
		}
	}

	for _, base := range r.ImplicitBases {
		if r.Literal(base, needle) {
			return pointer.NewResolution(constant.VerdictLive)
		}
	}

	return pointer.NewResolution(constant.VerdictDead)
}
