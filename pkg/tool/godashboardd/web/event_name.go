package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
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

	return join.Empty(constant.RowEventPrefix, slug)
}
