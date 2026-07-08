package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
	"unicode"
)

func IsSymbol(s string) bool {
	base := s[strings.LastIndex(s, separator.Slash)+1:]

	if base == "" {
		return false
	}

	if strings.HasSuffix(base, "()") {
		return true
	}

	if i := strings.LastIndex(base, separator.Dot); i != -1 {
		remainder := base[i+1:]

		return remainder != "" && unicode.IsUpper(rune(remainder[0]))
	}

	return unicode.IsUpper(rune(base[0]))
}
