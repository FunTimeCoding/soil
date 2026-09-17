package pointer

import "unicode"

func isPathRune(r rune) bool {
	return unicode.IsLetter(r) ||
		unicode.IsDigit(r) ||
		r == '_' ||
		r == '-' ||
		r == '.' ||
		r == '/'
}
