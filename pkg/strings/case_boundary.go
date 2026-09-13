package strings

import "unicode"

func caseBoundary(
	text []rune,
	index int,
) bool {
	previous := text[index-1]

	if unicode.IsLower(previous) || unicode.IsDigit(previous) {
		return true
	}

	return unicode.IsUpper(previous) &&
		index+1 < len(text) &&
		unicode.IsLower(text[index+1])
}
