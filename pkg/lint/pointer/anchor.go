package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	constant1 "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Anchor(
	base string,
	candidate string,
) (string, string) {
	n := Normalize(candidate)

	if n == "" ||
		strings.HasPrefix(n, constant1.Slash) ||
		strings.HasPrefix(n, constant.HomePrefix) {
		return "", ""
	}

	segment, _, _ := strings.Cut(n, constant1.Slash)

	return join.Empty(base, constant1.Slash, segment),
		join.Empty(base, constant1.Slash, n)
}
