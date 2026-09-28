package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func placementTable(
	v []*placement.Placement,
	place bool,
	empty string,
) gomponents.Node {
	if len(v) == 0 {
		return html.P(html.Em(gomponents.Text(empty)))
	}

	scoped := AnyScope(v)

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text(constant.NameHeader)),
				gomponents.If(
					scoped,
					html.Th(gomponents.Text(constant.ScopeHeader)),
				),
				gomponents.If(
					place,
					html.Th(gomponents.Text(constant.PlaceHeader)),
				),
				html.Th(gomponents.Text(constant.SourceHeader)),
				html.Th(gomponents.Text(constant.SeenHeader)),
			),
		),
		html.TBody(
			gomponents.Map(
				v,
				func(p *placement.Placement) gomponents.Node {
					return html.Tr(
						html.Td(gomponents.Text(p.Name)),
						gomponents.If(
							scoped,
							html.Td(gomponents.Text(p.Scope)),
						),
						gomponents.If(
							place,
							placeCell(p.PlaceKind, p.PlaceName),
						),
						html.Td(html.Small(gomponents.Text(p.Source))),
						layout.TimeCell(p.SeenAt),
					)
				},
			),
		),
	)
}
