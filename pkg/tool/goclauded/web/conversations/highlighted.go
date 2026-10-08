package conversations

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func highlighted(
	text string,
	terms []string,
) []gomponents.Node {
	lower := strings.ToLower(text)

	if len(lower) != len(text) || len(terms) == 0 {
		return []gomponents.Node{gomponents.Text(text)}
	}

	var lowered []string

	for _, term := range terms {
		lowered = append(lowered, strings.ToLower(term))
	}

	var result []gomponents.Node
	position := 0

	for position < len(text) {
		start, length := nextMatch(lower, lowered, position)

		if start < 0 {
			break
		}

		if start > position {
			result = append(result, gomponents.Text(text[position:start]))
		}

		result = append(
			result,
			html.Mark(gomponents.Text(text[start:start+length])),
		)
		position = start + length
	}

	if position < len(text) {
		result = append(result, gomponents.Text(text[position:]))
	}

	return result
}
