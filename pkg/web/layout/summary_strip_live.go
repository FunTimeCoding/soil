package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func summaryStripLive(items []string) gomponents.Node {
	return html.Div(
		extended.StreamSwap(constant.LayoutSummaryStrip),
		SummaryStripContent(items),
	)
}
