package highlight

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func (h *Highlight) Node() gomponents.Node {
	if h.open != "" {
		return h.marked()
	}

	lower := strings.ToLower(h.text)

	if len(lower) != len(h.text) || len(h.terms) == 0 {
		return gomponents.Text(h.text)
	}

	var lowered []string

	for _, term := range h.terms {
		lowered = append(lowered, strings.ToLower(term))
	}

	var result []gomponents.Node
	position := 0

	for position < len(h.text) {
		start, length := nextMatch(lower, lowered, position)

		if start < 0 {
			break
		}

		if start > position {
			result = append(result, gomponents.Text(h.text[position:start]))
		}

		result = append(
			result,
			html.Mark(
				html.Class(constant.HighlightClass),
				gomponents.Text(h.text[start:start+length]),
			),
		)
		position = start + length
	}

	if position < len(h.text) {
		result = append(result, gomponents.Text(h.text[position:]))
	}

	return gomponents.Group(result)
}
