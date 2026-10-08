package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"sort"
	"strings"
)

func Locate(
	content string,
	blocks []*block.Block,
) ([]*Match, []string) {
	var result []*Match
	var failures []string

	for _, b := range blocks {
		offsets := occurrences(content, b.Search)

		switch len(offsets) {
		case 1:
			result = append(
				result,
				&Match{
					Block:  b,
					Offset: offsets[0],
					Line:   lineAt(content, offsets[0]),
				},
			)
		case 0:
			failures = append(failures, notFound(content, b))
		default:
			var lines []string

			for _, o := range offsets {
				lines = append(lines, fmt.Sprint(lineAt(content, o)))
			}

			failures = append(
				failures,
				fmt.Sprintf(
					constant.Ambiguous,
					b.Number,
					len(offsets),
					strings.Join(lines, constant.LineSeparator),
				),
			)
		}
	}

	sort.Slice(
		result,
		func(i, j int) bool { return result[i].Offset < result[j].Offset },
	)

	for i := 1; i < len(result); i++ {
		previous := result[i-1]

		if previous.Offset+len(previous.Block.Search) > result[i].Offset {
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

	return result, failures
}
