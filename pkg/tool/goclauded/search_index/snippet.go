package search_index

import (
	soil "github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
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

	from := soil.AlignRuneStart(text, max(0, start-constant.SnippetRadius))
	to := soil.AlignRuneStart(
		text,
		min(len(text), start+length+constant.SnippetRadius),
	)
	result := text[from:to]

	if from > 0 {
		result = join.Empty("…", result)
	}

	if to < len(text) {
		result = join.Empty(result, "…")
	}

	return result
}
