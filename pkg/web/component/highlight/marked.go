package highlight

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func (h *Highlight) marked() gomponents.Node {
	var result []gomponents.Node
	rest := h.text

	for {
		before, inside, found := strings.Cut(rest, h.open)

		if !found {
			break
		}

		term, after, closed := strings.Cut(inside, h.close)

		if !closed {
			break
		}

		if before != "" {
			result = append(result, gomponents.Text(before))
		}

		result = append(
			result,
			html.Mark(
				html.Class(constant.HighlightClass),
				gomponents.Text(term),
			),
		)
		rest = after
	}

	result = append(
		result,
		gomponents.Text(
			strings.ReplaceAll(
				strings.ReplaceAll(rest, h.open, ""),
				h.close,
				"",
			),
		),
	)

	return gomponents.Group(result)
}
