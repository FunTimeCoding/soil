package heading

import (
	"strings"
	"unicode"
)

func Slug(text string) string {
	var result strings.Builder

	for _, r := range strings.ToLower(text) {
		switch {
		case r == ' ':
			result.WriteRune('-')
		case r == '-' ||
			unicode.IsLetter(r) ||
			unicode.IsNumber(r) ||
			unicode.IsMark(r) ||
			unicode.Is(unicode.Pc, r):
			result.WriteRune(r)
		}
	}

	return result.String()
}
