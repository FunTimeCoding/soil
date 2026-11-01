package summary_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_row"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(
	rows []summary_row.Row,
	group string,
) gomponents.Node {
	surface := group == constant.Surface
	kind := group == constant.Kind
	columns := []string{"Tool"}

	if surface {
		columns = append(columns, "Surface")
	}

	if kind {
		columns = append(columns, "Kind")
	}

	columns = append(columns, "Count")
	nodes := make([]gomponents.Node, 0, len(rows))

	for _, r := range rows {
		nodes = append(
			nodes,
			html.Tr(
				html.Td(gomponents.Text(r.Tool)),
				gomponents.If(surface, html.Td(gomponents.Text(r.Surface))),
				gomponents.If(kind, html.Td(gomponents.Text(r.Kind))),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Text(fmt.Sprintf("%d", r.Count)),
				),
			),
		)
	}

	return table.New(columns...).
		Numeric(len(columns) - 1).
		Rows(nodes).
		Empty("No events recorded.").
		Node()
}
