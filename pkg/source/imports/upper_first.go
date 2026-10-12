package imports

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"unicode"
	"unicode/utf8"
)

func upperFirst(name string) string {
	r, size := utf8.DecodeRuneInString(name)

	return join.Empty(string(unicode.ToUpper(r)), name[size:])
}
