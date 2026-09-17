package pointer

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Anchor(
	base string,
	candidate string,
) (string, string) {
	n := Normalize(candidate)

	if n == "" ||
		strings.HasPrefix(n, constant.Slash) ||
		strings.HasPrefix(n, lintConstant.HomePrefix) {
		return "", ""
	}

	segment, _, _ := strings.Cut(n, constant.Slash)

	return join.Empty(base, constant.Slash, segment),
		join.Empty(base, constant.Slash, n)
}
