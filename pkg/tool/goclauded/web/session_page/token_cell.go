package session_page

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func tokenCell(count int) gomponents.Node {
	return html.Td(
		html.Class(webConstant.NumericClass),
		gomponents.Text(pricing.FormatTokens(count)),
	)
}
