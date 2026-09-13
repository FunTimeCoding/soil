package dictionary

import (
	"strings"
	"unicode"
)

func splitDigit(line string) string {
	text := []rune(line)
	var b strings.Builder

	for i, current := range text {
		if i > 0 && unicode.IsDigit(current) != unicode.IsDigit(text[i-1]) {
			b.WriteRune(' ')
		}

		b.WriteRune(current)
	}

	return b.String()
}
