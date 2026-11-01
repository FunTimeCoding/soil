package log_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(fights []fight.Fight) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(fights))

	for _, f := range fights {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.Input(
						html.Type(webConstant.EntryCheckboxKind),
						html.Name(constant.FileNameField),
						html.Value(f.Filename),
					),
				),
				layout.TimeCell(f.Timestamp),
				html.Td(gomponents.Text(length(f))),
				html.Td(gomponents.Text(mapName(f))),
				component.NumericCell(fmt.Sprintf("%d", f.AlliedCount)),
				component.NumericCell(enemies(f)),
			),
		)
	}

	return table.New(
		"",
		"Time",
		"Duration",
		"Map",
		constant.PlayersTitle,
		"Enemies",
	).
		Numeric(4, 5).
		Rows(rows).
		Empty("No logs found.").
		Node()
}
