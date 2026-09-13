package scan

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"regexp"
	"strings"
)

func New(terms []string) *Scanner {
	result := &Scanner{}
	seen := map[string]bool{}

	for _, term := range terms {
		if term == "" {
			continue
		}

		key := strings.ToLower(term)

		if seen[key] {
			continue
		}

		seen[key] = true
		result.terms = append(result.terms, term)
		result.lower = append(result.lower, key)
		result.patterns = append(
			result.patterns,
			regexp.MustCompile(
				join.Empty(`(?i)\b`, regexp.QuoteMeta(term), `\b`),
			),
		)
	}

	return result
}
