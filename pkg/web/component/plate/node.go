package plate

import (
	"github.com/funtimecoding/soil/pkg/web/component/fact/row"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Plate) Node() gomponents.Node {
	return html.Div(
		html.Class(constant.PlateClass),
		line(constant.PlateNameClass, p.name),
		line(constant.PlateHeadlineClass, p.headline),
		line(constant.PlateCaptionClass, p.caption),
		gomponents.Map(
			p.rows,
			func(r *row.Row) gomponents.Node {
				return html.Div(
					html.Class(constant.PlateRowClass),
					html.Span(
						html.Class(constant.PlateLabelClass),
						gomponents.Text(r.Label),
					),
					html.Div(html.Class(constant.PlateValueClass), r.Value),
				)
			},
		),
	)
}
