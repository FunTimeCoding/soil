package layout

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func summaryStrip(items []string) gomponents.Node {
	return html.Div(SummaryStripContent(items))
}
