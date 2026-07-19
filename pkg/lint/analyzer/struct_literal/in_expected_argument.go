package struct_literal

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
	"go/token"
)

func inExpectedArgument(
	ranges []assert_call.Range,
	position token.Pos,
) bool {
	for _, r := range ranges {
		if position >= r.From && position < r.To {
			return true
		}
	}

	return false
}
