package component

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func PlacementTable(
	v []*placement.Placement,
	place bool,
	empty string,
) gomponents.Node {
	scoped := AnyScope(v)
	columns := []string{constant.NameHeader}

	if scoped {
		columns = append(columns, constant.ScopeHeader)
	}

	if place {
		columns = append(columns, constant.PlaceHeader)
	}

	columns = append(columns, constant.SourceHeader, constant.SeenHeader)
	rows := make([]gomponents.Node, 0, len(v))

	for _, p := range v {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(p.Name)),
				gomponents.If(scoped, html.Td(gomponents.Text(p.Scope))),
				gomponents.If(place, PlaceCell(p.PlaceKind, p.PlaceName)),
				html.Td(html.Small(gomponents.Text(p.Source))),
				layout.TimeCell(p.SeenAt),
			),
		)
	}

	return table.New(columns...).Rows(rows).Empty(empty).Node()
}
