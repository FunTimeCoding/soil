package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Anchor(
	base string,
	candidate string,
) (string, string) {
	n := Normalize(candidate)

	if n == "" ||
		strings.HasPrefix(n, stringsConstant.Slash) ||
		strings.HasPrefix(n, constant.HomePrefix) {
		return "", ""
	}

	segment, _, _ := strings.Cut(n, stringsConstant.Slash)

	return join.Empty(base, stringsConstant.Slash, segment),
		join.Empty(base, stringsConstant.Slash, n)
}
