package camel

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"unicode"
)

func FromSnake(s string) string {
	var b strings.Builder

	for _, word := range strings.Split(s, constant.Underscore) {
		if word == "" {
			continue
		}

		r := []rune(word)
		b.WriteRune(unicode.ToUpper(r[0]))
		b.WriteString(string(r[1:]))
	}

	return b.String()
}
