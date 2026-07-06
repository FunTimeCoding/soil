package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func parseQuotedPhrase(
	s string,
	i *int,
) string {
	*i++
	start := *i

	for *i < len(s) && s[*i] != '"' {
		*i++
	}

	phrase := strings.TrimSpace(s[start:*i])

	if *i < len(s) {
		*i++
	}

	words := strings.Fields(phrase)
	var sanitized []string

	for _, w := range words {
		t := sanitizeTerm(w)

		if t != "" {
			sanitized = append(sanitized, t)
		}
	}

	return join.Space(sanitized...)
}
