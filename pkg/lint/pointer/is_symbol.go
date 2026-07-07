package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
	"unicode"
)

func IsSymbol(s string) bool {
	base := s[strings.LastIndex(s, separator.Slash)+1:]

	if base == "" || strings.Contains(base, separator.Dot) {
		return false
	}

	return unicode.IsUpper(rune(base[0]))
}
