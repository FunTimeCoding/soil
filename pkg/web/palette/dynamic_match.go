package palette

import "github.com/funtimecoding/soil/pkg/web/constant"

func dynamicMatch(
	pattern []rune,
	text []rune,
	bonuses []int,
) (int, []int) {
	rows := len(pattern)
	columns := len(text)
	scores := make([][]int, rows)
	consecutive := make([][]int, rows)

	for i := range rows {
		scores[i] = make([]int, columns)
		consecutive[i] = make([]int, columns)

		for j := range columns {
			scores[i][j] = constant.PaletteMinScore
		}
	}

	for j := range columns {
		if text[j] != pattern[0] {
			continue
		}

		bonus := bonuses[j] * constant.PaletteBonusFirstCharacter
		scores[0][j] = constant.PaletteScoreMatch + bonus
		consecutive[0][j] = 1
	}

	for i := 1; i < rows; i++ {
		for j := i; j < columns; j++ {
			if text[j] != pattern[i] {
				continue
			}

			bonus := bonuses[j]
			diagonalScore := constant.PaletteMinScore

			if j > 0 && scores[i-1][j-1] > constant.PaletteMinScore {
				past := consecutive[i-1][j-1]
				consecutiveBonus := max(
					bonus,
					constant.PaletteBonusConsecutiveMin,
				)

				if past > 0 {
					bonus = consecutiveBonus
				}

				diagonalScore = scores[i-1][j-1] + constant.PaletteScoreMatch + bonus
			} else if i == 0 {
				diagonalScore = constant.PaletteScoreMatch + bonus*constant.PaletteBonusFirstCharacter
			}

			gapScore := constant.PaletteMinScore

			for k := i - 1; k < j; k++ {
				if scores[i-1][k] <= constant.PaletteMinScore {
					continue
				}

				gap := j - k - 1
				penalty := constant.PaletteScoreGapStart + constant.PaletteScoreGapExtension*gap
				candidate := scores[i-1][k] + penalty + constant.PaletteScoreMatch + bonus

				if candidate > gapScore {
					gapScore = candidate
				}
			}

			best := max(diagonalScore, gapScore)

			if best > constant.PaletteMinScore {
				scores[i][j] = best

				if diagonalScore >= gapScore && j > 0 {
					consecutive[i][j] = consecutive[i-1][j-1] + 1
				} else {
					consecutive[i][j] = 1
				}
			}
		}
	}

	bestScore := constant.PaletteMinScore
	bestEnd := -1

	for j := rows - 1; j < columns; j++ {
		if scores[rows-1][j] > bestScore {
			bestScore = scores[rows-1][j]
			bestEnd = j
		}
	}

	if bestEnd < 0 {
		return -1, nil
	}

	positions := backtrack(scores, consecutive, pattern, text, bestEnd)

	return bestScore, positions
}
