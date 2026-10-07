package pointer

import "unicode"

func IsPathRune(r rune) bool {
	return unicode.IsLetter(r) ||
		unicode.IsDigit(r) ||
		r == '_' ||
		r == '-' ||
		r == '.' ||
		r == '/'
}
