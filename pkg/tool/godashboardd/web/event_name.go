package web

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"strings"
)

func eventName(label string) string {
	slug := strings.Map(
		func(r rune) rune {
			if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
				return r
			}

			if r >= 'A' && r <= 'Z' {
				return r + 'a' - 'A'
			}

			return '-'
		},
		label,
	)

	return join.Empty(rowEventPrefix, slug)
}
