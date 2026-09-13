package strings

import (
	"strings"
	"unicode"
)

func SplitCase(line string) string {
	text := []rune(line)
	var b strings.Builder

	for i, current := range text {
		if i > 0 && unicode.IsUpper(current) && caseBoundary(text, i) {
			b.WriteRune(' ')
		}

		b.WriteRune(current)
	}

	return b.String()
}
