package raid_page

import (
	"fmt"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(raids []record.Raid) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(raids))

	for _, r := range raids {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(component.RaidPath(r.Identifier)),
						gomponents.Text(r.Name),
					),
				),
				html.Td(
					html.Class(webConstant.LayoutTimeCellClass),
					gomponents.Text(r.Date.Format(timeConstant.DateYear)),
				),
				component.NumericCell(fmt.Sprintf("%d", r.Fights)),
				component.NumericCell(fmt.Sprintf("%d", r.Players)),
			),
		)
	}

	return table.New("Name", "Date", "Fights", constant.PlayersTitle).
		Numeric(2, 3).
		Rows(rows).
		Empty("No raids created yet.").
		Node()
}
