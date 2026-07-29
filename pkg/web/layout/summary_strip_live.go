package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func summaryStripLive(items []string) gomponents.Node {
	return html.Div(
		gomponents.Attr("sse-swap", constant.LayoutSummaryStrip),
		SummaryStripContent(items),
	)
}
