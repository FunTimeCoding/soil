package markup_tester

import (
	"regexp"
	"strings"
)

func Alternatives(selector string) []string {
	s := regexp.MustCompile(
		`:(focus-visible|focus-within|hover|focus|active|visited)`,
	).ReplaceAllString(selector, "")
	s = strings.ReplaceAll(s, ":has(> ", ":has(")

	return expand(s)
}
