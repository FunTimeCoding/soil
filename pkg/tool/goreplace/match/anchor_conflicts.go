package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
)

func anchorConflicts(matches []*Match) []string {
	var result []string

	for _, insert := range matches {
		if insert.Length > 0 {
			continue
		}

		anchorEnd := insert.Anchor + len(insert.Block.Search)

		for _, other := range matches {
			if other.Length == 0 ||
				insert.Anchor >= other.Offset+other.Length ||
				anchorEnd <= other.Offset {
				continue
			}

			result = append(
				result,
				fmt.Sprintf(
					constant.AnchorRemoved,
					insert.Block.Number,
					insert.Line,
					other.Block.Number,
				),
			)
		}
	}

	return result
}
