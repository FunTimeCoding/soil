package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
)

func rewrite(
	r *resolver.Resolver,
	source string,
	declared *pointer.Declared,
	moves *move_set.Set,
	c *pointer.Candidate,
) (string, string) {
	if pointer.IsTestName(c.Span) {
		after, unique := testSpan(r, moves, c.Span)

		if !unique {
			return "", constant.RepeatedTestReason
		}

		return after, ""
	}

	after, rewritable := rewriteSpan(r, source, declared, moves, c)

	if !rewritable {
		return "", constant.BracedSpanReason
	}

	return after, ""
}
