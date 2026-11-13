package fact

import (
	"github.com/funtimecoding/soil/pkg/web/component/fact/row"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) table() gomponents.Node {
	return html.Table(
		html.Class(constant.FactsClass),
		html.TBody(
			gomponents.Map(
				f.rows,
				func(r *row.Row) gomponents.Node {
					if r.Section {
						return html.Tr(
							html.Class(constant.SectionClass),
							html.Th(html.ColSpan("2"), gomponents.Text(r.Label)),
						)
					}

					return html.Tr(
						html.Th(
							html.Scope(constant.RowScope),
							gomponents.Text(r.Label),
						),
						html.Td(
							gomponents.If(r.Tooltip == "", r.Value),
							gomponents.If(
								r.Tooltip != "",
								html.Span(
									gomponents.Attr(
										constant.TooltipAttribute,
										r.Tooltip,
									),
									r.Value,
								),
							),
						),
					)
				},
			)...,
		),
	)
}
