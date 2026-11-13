package fact

import (
	"github.com/funtimecoding/soil/pkg/web/component/fact/row"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func item(r *row.Row) gomponents.Node {
	return html.Div(
		html.Div(
			html.Class(constant.GlanceLabelClass),
			gomponents.Text(r.Label),
		),
		html.Div(
			html.Class(constant.GlanceValueClass),
			gomponents.If(
				r.Tooltip != "",
				gomponents.Attr(constant.TooltipAttribute, r.Tooltip),
			),
			r.Value,
		),
	)
}
