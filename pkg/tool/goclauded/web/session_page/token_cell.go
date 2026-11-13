package session_page

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func tokenCell(count int) gomponents.Node {
	return html.Td(gomponents.Text(pricing.FormatTokens(count)))
}
