package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sightingTable(
	v []*sighting.Sighting,
	empty string,
) gomponents.Node {
	if len(v) == 0 {
		return html.P(html.Em(gomponents.Text(empty)))
	}

	placed := AnyPlace(v)

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text(constant.AddressHeader)),
				html.Th(gomponents.Text(constant.HardwareHeader)),
				html.Th(gomponents.Text(constant.HostnameHeader)),
				gomponents.If(
					placed,
					html.Th(gomponents.Text(constant.PlaceHeader)),
				),
				html.Th(gomponents.Text(constant.SeenHeader)),
			),
		),
		html.TBody(
			gomponents.Map(
				v,
				func(e *sighting.Sighting) gomponents.Node {
					return html.Tr(
						html.Td(gomponents.Text(e.Address)),
						html.Td(html.Small(gomponents.Text(e.HardwareAddress))),
						html.Td(gomponents.Text(e.Hostname)),
						gomponents.If(
							placed,
							placeCell(e.PlaceKind, e.PlaceName),
						),
						layout.TimeCell(e.SeenAt),
					)
				},
			),
		),
	)
}
