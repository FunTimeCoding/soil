package text

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"regexp"
	"strings"
)

func SpaceToUnderscore(s string) string {
	return strings.TrimSpace(
		regexp.MustCompile(`\s+`).ReplaceAllString(s, constant.Underscore),
	)
}
