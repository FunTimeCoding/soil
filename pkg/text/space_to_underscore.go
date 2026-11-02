package text

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"regexp"
	"strings"
)

func SpaceToUnderscore(s string) string {
	return strings.TrimSpace(
		regexp.MustCompile(`\s+`).ReplaceAllString(
			s,
			stringsConstant.Underscore,
		),
	)
}
