package palette

import "github.com/funtimecoding/soil/pkg/web/constant"

func characterBonus(
	previous rune,
	current rune,
) int {
	if isWhitespace(previous) {
		return constant.PaletteBonusWhitespace
	}

	if isLower(previous) && isUpper(current) {
		return constant.PaletteBonusCamelCase
	}

	if isDelimiter(previous) {
		return constant.PaletteBonusWordBoundary
	}

	return 0
}
