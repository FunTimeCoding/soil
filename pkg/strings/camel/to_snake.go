package camel

import (
	"strings"
	"unicode"
)

func ToSnake(s string) string {
	var b strings.Builder

	for i, r := range s {
		if i > 0 && unicode.IsDigit(r) && unicode.IsLower(rune(s[i-1])) {
			b.WriteByte('_')
		}

		if unicode.IsUpper(r) {
			if i > 0 {
				previous := rune(s[i-1])

				if unicode.IsLower(previous) || unicode.IsDigit(previous) {
					b.WriteByte('_')
				} else if unicode.IsUpper(previous) && i+1 < len(s) && unicode.IsLower(rune(s[i+1])) {
					b.WriteByte('_')
				}
			}

			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}
