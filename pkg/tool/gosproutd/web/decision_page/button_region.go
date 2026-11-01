package decision_page

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func buttonRegion(
	identifier uint,
	part string,
	content gomponents.Node,
) gomponents.Node {
	return html.Div(
		html.Class(webConstant.ClusterClass),
		extended.StreamSwap(Event(identifier, part)),
		content,
	)
}
