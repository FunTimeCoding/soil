package service

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func FlipName(name string) string {
	r, size := utf8.DecodeRuneInString(name)

	if unicode.IsUpper(r) {
		return fmt.Sprintf("%c%s", unicode.ToLower(r), name[size:])
	}

	return fmt.Sprintf("%c%s", unicode.ToUpper(r), name[size:])
}
