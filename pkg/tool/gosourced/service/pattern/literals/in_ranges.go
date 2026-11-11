package literals

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
	"go/token"
)

func inRanges(
	ranges []assert_call.Range,
	position token.Pos,
) bool {
	for _, r := range ranges {
		if position >= r.Start && position < r.End {
			return true
		}
	}

	return false
}
