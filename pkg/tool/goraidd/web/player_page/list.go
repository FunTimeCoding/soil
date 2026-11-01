package player_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(rows []record.Attendance) gomponents.Node {
	nodes := make([]gomponents.Node, 0, len(rows))

	for _, r := range rows {
		share := 0

		if r.Available > 0 {
			share = r.Fights * 100 / r.Available
		}

		nodes = append(
			nodes,
			html.Tr(
				html.Td(
					html.A(
						html.Href(component.PlayerPath(r.Account)),
						gomponents.Text(r.Account),
					),
				),
				html.Td(gomponents.Text(r.Characters)),
				component.NumericCell(
					fmt.Sprintf("%d / %d", r.Fights, r.Available),
				),
				component.NumericCell(fmt.Sprintf("%d%%", share)),
			),
		)
	}

	return table.New("Account", "Characters", "Fights", "Attendance").
		Numeric(2, 3).
		Rows(nodes).
		Empty("No player data yet.").
		Node()
}
