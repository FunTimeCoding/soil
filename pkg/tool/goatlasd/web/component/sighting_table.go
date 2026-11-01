package component

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SightingTable(
	v []*sighting.Sighting,
	empty string,
) gomponents.Node {
	placed := AnyPlace(v)
	columns := []string{
		constant.AddressHeader,
		constant.HardwareHeader,
		constant.HostnameHeader,
	}

	if placed {
		columns = append(columns, constant.PlaceHeader)
	}

	columns = append(columns, constant.SeenHeader)
	rows := make([]gomponents.Node, 0, len(v))

	for _, e := range v {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(e.Address)),
				html.Td(html.Small(gomponents.Text(e.HardwareAddress))),
				html.Td(gomponents.Text(e.Hostname)),
				gomponents.If(placed, PlaceCell(e.PlaceKind, e.PlaceName)),
				layout.TimeCell(e.SeenAt),
			),
		)
	}

	return table.New(columns...).Rows(rows).Empty(empty).Node()
}
