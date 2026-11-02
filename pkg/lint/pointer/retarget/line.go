package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/retarget_change"
	"strings"
)

func Line(
	r *resolver.Resolver,
	source string,
	declared *pointer.Declared,
	moves *move_set.Set,
	text string,
	number int,
) (string, []*retarget_change.Change) {
	result := text
	var changes []*retarget_change.Change

	for _, c := range pointer.Extract(text) {
		after, rewritable := rewriteSpan(r, source, declared, moves, c)

		if !rewritable {
			changes = append(
				changes,
				retarget_change.NewManual(
					number,
					c.Span,
					"braced span needs a hand edit",
				),
			)

			continue
		}

		if after == "" {
			continue
		}

		before := wrap(c.Link, c.Span)

		if !strings.Contains(result, before) {
			changes = append(
				changes,
				retarget_change.NewManual(
					number,
					c.Span,
					"span is not written verbatim",
				),
			)

			continue
		}

		result = strings.Replace(result, before, wrap(c.Link, after), 1)
		changes = append(changes, retarget_change.New(number, c.Span, after))
	}

	return result, changes
}
