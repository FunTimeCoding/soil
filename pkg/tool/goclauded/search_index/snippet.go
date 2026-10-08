package search_index

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
	"unicode/utf8"
)

func snippet(
	body string,
	terms []string,
) string {
	text := strings.Join(strings.Fields(body), " ")
	lower := strings.ToLower(text)
	start := -1
	length := 0

	for _, term := range terms {
		i := strings.Index(lower, strings.ToLower(term))

		if i >= 0 && (start < 0 || i < start) {
			start = i
			length = len(term)
		}
	}

	if start < 0 || len(lower) != len(text) {
		start = 0
	}

	from := max(0, start-constant.SnippetRadius)
	to := min(len(text), start+length+constant.SnippetRadius)

	for from > 0 && !utf8.RuneStart(text[from]) {
		from--
	}

	for to < len(text) && !utf8.RuneStart(text[to]) {
		to++
	}

	result := text[from:to]

	if from > 0 {
		result = join.Empty("…", result)
	}

	if to < len(text) {
		result = join.Empty(result, "…")
	}

	return result
}
