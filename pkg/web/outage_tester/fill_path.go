package outage_tester

import (
	"regexp"
	"strings"
)

func fillPath(
	path string,
	parameters map[string]string,
) string {
	return regexp.MustCompile(`\{[^}]+\}`).ReplaceAllStringFunc(
		path,
		func(p string) string {
			if v, okay := parameters[strings.Trim(p, "{}")]; okay {
				return v
			}

			return "1"
		},
	)
}
