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
		return pointer.NewUnchecked(constant.ReasonRoute)
	}

	needle := pointer.RouteLiteral(route)

	for _, base := range bases {
		if strings.HasPrefix(route, constant.RestRoutePrefix) {
			if paths, found := r.Routes(base); found && pointer.MatchRoute(
				route,
				paths,
			) {
				return pointer.NewLive("")
			}
		}

		if r.Literal(base, needle) {
			return pointer.NewLive("")
		}
	}

	for _, base := range r.ImplicitBases {
		if r.Literal(base, needle) {
			return pointer.NewLive("")
		}
	}

	return pointer.NewConcern(constant.DeadPointerKey)
}
