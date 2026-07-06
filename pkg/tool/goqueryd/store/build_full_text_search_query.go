package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func buildFullTextSearchQuery(query string) string {
	s := strings.TrimSpace(query)
	var positive, negative []string
	i := 0

	for i < len(s) {
		for i < len(s) && isWhitespace(s[i]) {
			i++
		}

		if i >= len(s) {
			break
		}

		negated := s[i] == '-'

		if negated {
			i++
		}

		if i >= len(s) {
			break
		}

		if s[i] == '"' {
			term := parseQuotedPhrase(s, &i)

			if term != "" {
				if negated {
					negative = append(negative, join.Empty(`"`, term, `"`))
				} else {
					positive = append(positive, join.Empty(`"`, term, `"`))
				}
			}
		} else {
			raw := parsePlainTerm(s, &i)

			if isHyphenatedToken(raw) {
				term := sanitizeHyphenatedTerm(raw)

				if term != "" {
					if negated {
						negative = append(
							negative,
							join.Empty(`"`, term, `"`),
						)
					} else {
						positive = append(
							positive,
							join.Empty(`"`, term, `"`),
						)
					}
				}
			} else {
				term := sanitizeTerm(raw)

				if term != "" {
					if negated {
						negative = append(
							negative,
							join.Empty(`"`, term, `"*`),
						)
					} else {
						positive = append(
							positive,
							join.Empty(`"`, term, `"*`),
						)
					}
				}
			}
		}
	}

	if len(positive) == 0 {
		return ""
	}

	result := strings.Join(positive, " AND ")

	for _, n := range negative {
		result = join.Empty(result, " NOT ", n)
	}

	return result
}
