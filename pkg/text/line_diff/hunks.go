package line_diff

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/text/constant"
)

func Hunks(
	lines []*Line,
	context int,
) []string {
	shown := make([]bool, len(lines))

	for i, l := range lines {
		if l.Mark == constant.DiffEqual {
			continue
		}

		for j := i - context; j <= i+context; j++ {
			if j >= 0 && j < len(lines) {
				shown[j] = true
			}
		}
	}

	var result []string
	gap := false

	for i, l := range lines {
		if !shown[i] {
			gap = true

			continue
		}

		if gap {
			result = append(result, constant.DiffGap)
			gap = false
		}

		result = append(result, join.Empty(l.Mark, l.Text))
	}

	return result
}
