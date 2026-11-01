package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func rewriteSpan(
	r *resolver.Resolver,
	source string,
	declared *pointer.Declared,
	moves []*path_move.Move,
	c *pointer.Candidate,
) (string, bool) {
	alternatives := pointer.Expand(c)

	if len(alternatives) == 1 {
		return alternative(r, source, declared, moves, c), true
	}

	var afters []string
	changed := false

	for _, a := range alternatives {
		after := alternative(r, source, declared, moves, a)

		if after == "" {
			after = a.Span
		} else {
			changed = true
		}

		afters = append(afters, after)
	}

	if !changed {
		return "", true
	}

	start := strings.Index(c.Span, "{")
	stop := start + strings.Index(c.Span[start:], "}")
	inner := c.Span[start+1 : stop]
	post := c.Span[stop+1:]
	choices := strings.Split(inner, constant.Comma)

	if len(choices) != len(afters) {
		return "", false
	}

	head := ""

	for i, after := range afters {
		tail := join.Empty(choices[i], post)

		if !strings.HasSuffix(after, tail) {
			return "", false
		}

		h := strings.TrimSuffix(after, tail)

		if i > 0 && h != head {
			return "", false
		}

		head = h
	}

	return join.Empty(head, "{", inner, "}", post), true
}
