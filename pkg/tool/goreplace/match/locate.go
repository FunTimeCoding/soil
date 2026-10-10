package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"sort"
)

func Locate(
	content string,
	blocks []*block.Block,
) ([]*Match, []string) {
	var result []*Match
	var failures []string

	for _, b := range blocks {
		m, failure := region(content, b)

		if failure != "" {
			failures = append(failures, failure)

			continue
		}

		result = append(result, m)
	}

	sort.SliceStable(
		result,
		func(i, j int) bool {
			if result[i].Offset != result[j].Offset {
				return result[i].Offset < result[j].Offset
			}

			return result[i].Length < result[j].Length
		},
	)

	for i := 1; i < len(result); i++ {
		previous := result[i-1]

		if previous.Offset+previous.Length > result[i].Offset {
			failures = append(
				failures,
				fmt.Sprintf(
					constant.Overlap,
					previous.Block.Number,
					result[i].Block.Number,
					result[i].Line,
				),
			)
		}
	}

	return result, append(failures, anchorConflicts(result)...)
}
