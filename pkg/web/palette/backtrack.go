package palette

import "github.com/funtimecoding/soil/pkg/web/constant"

func backtrack(
	scores [][]int,
	consecutive [][]int,
	pattern []rune,
	text []rune,
	endColumn int,
) []int {
	positions := make([]int, len(pattern))
	j := endColumn

	for i := len(pattern) - 1; i >= 0; i-- {
		positions[i] = j

		if i > 0 {
			if consecutive[i][j] > 1 && j > 0 {
				j--
			} else {
				best := constant.PaletteMinScore
				bestK := -1

				for k := range j {
					if scores[i-1][k] > best {
						best = scores[i-1][k]
						bestK = k
					}
				}

				j = bestK
			}
		}
	}

	return positions
}
