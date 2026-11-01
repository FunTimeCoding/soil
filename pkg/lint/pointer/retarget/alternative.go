package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func alternative(
	r *resolver.Resolver,
	source string,
	declared *pointer.Declared,
	moves []*path_move.Move,
	c *pointer.Candidate,
) string {
	span, fragment, cut := strings.Cut(c.Span, constant.FragmentSeparator)

	if span == "" {
		return ""
	}

	after := ""

	if rest, symbol := strings.CutPrefix(span, constant.SchemeGo); symbol {
		after = symbolSpan(moves, rest)
	} else if movable(r, c.Link, span) {
		inner := pointer.NewSpan(span)
		inner.Link = c.Link
		v := r.Resolve(source, declared, inner)

		if v.Verdict != constant.VerdictLive || v.Target == "" {
			return ""
		}

		normalized := pointer.Normalize(span)
		path := rewritePath(moves, source, normalized, v.Target)

		if path != "" {
			after = strings.Replace(span, normalized, path, 1)
		}
	}

	if after == "" || !cut {
		return after
	}

	return join.Empty(after, constant.FragmentSeparator, fragment)
}
