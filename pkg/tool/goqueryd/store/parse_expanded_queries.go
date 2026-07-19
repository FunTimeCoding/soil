package store

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func ParseExpandedQueries(
	text string,
	originalQuery string,
) []ExpandedQuery {
	lower := strings.ToLower(originalQuery)
	var result []ExpandedQuery

	for _, line := range strings.Split(
		strings.TrimSpace(text),
		separator.Unix,
	) {
		colon := strings.Index(line, ":")

		if colon == -1 {
			continue
		}

		queryType := strings.TrimSpace(line[:colon])

		if queryType != "lex" && queryType != "vec" && queryType != "hyde" {
			continue
		}

		query := strings.TrimSpace(line[colon+1:])

		if query == "" || strings.ToLower(query) == lower {
			continue
		}

		result = append(
			result,
			ExpandedQuery{Type: queryType, Query: query},
		)
	}

	return result
}
