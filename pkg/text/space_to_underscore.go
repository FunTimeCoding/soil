package text

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"regexp"
	"strings"
)

func SpaceToUnderscore(s string) string {
	return strings.TrimSpace(
		regexp.MustCompile(`\s+`).ReplaceAllString(
			s,
			separator.Underscore,
		),
	)
}
