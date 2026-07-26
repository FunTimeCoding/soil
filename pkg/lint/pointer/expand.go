package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Expand(s string) []string {
	start := strings.Index(s, "{")

	if start == -1 {
		return []string{s}
	}

	stop := strings.Index(s[start:], "}")

	if stop == -1 {
		return []string{s}
	}

	inner := s[start+1 : start+stop]

	if !strings.Contains(inner, constant.Comma) {
		return []string{s}
	}

	var result []string

	for _, alternative := range strings.Split(inner, constant.Comma) {
		result = append(
			result,
			Expand(
				join.Empty(s[:start], alternative, s[start+stop+1:]),
			)...,
		)
	}

	return result
}
