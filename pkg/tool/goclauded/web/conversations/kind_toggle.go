package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func kindToggle(
	kind string,
	checked bool,
) gomponents.Node {
	return html.Label(
		html.Input(
			html.Type("checkbox"),
			html.Name(constant.KindField),
			html.Value(kind),
			gomponents.If(checked, html.Checked()),
		),
		gomponents.Text(kind),
	)
}
