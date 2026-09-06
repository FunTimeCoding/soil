package scan

import "strings"

func stripMethod(pattern string) string {
	before, after, found := strings.Cut(pattern, " ")

	if !found || before == "" {
		return pattern
	}

	for _, r := range before {
		if r < 'A' || r > 'Z' {
			return pattern
		}
	}

	return after
}
