package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func logsTable(
	fights []raid.Fight,
	offset, total int,
	startValue, endValue string,
	filtered bool,
) gomponents.Node {
	if len(fights) == 0 {
		return html.P(html.Em(gomponents.Text("No logs found.")))
	}

	var nodes []gomponents.Node

	if !filtered {
		nodes = append(
			nodes,
			paginationControls(offset, total, startValue, endValue),
		)
	}

	nodes = append(
		nodes,
		html.Table(
			html.THead(
				html.Tr(
					html.Th(),
					html.Th(gomponents.Text("Time")),
					html.Th(gomponents.Text("Duration")),
					html.Th(gomponents.Text("Map")),
					html.Th(gomponents.Text(constant.PlayersTitle)),
					html.Th(gomponents.Text("Enemies")),
				),
			),
			html.TBody(
				gomponents.Map(
					fights,
					func(f raid.Fight) gomponents.Node {
						return html.Tr(
							html.Td(
								html.Input(
									html.Type("checkbox"),
									html.Name("fileNames"),
									html.Value(f.Filename),
								),
							),
							html.Td(
								gomponents.Text(
									f.Timestamp.Format(time.DateMinute),
								),
							),
							html.Td(gomponents.Text(fightDuration(f))),
							html.Td(gomponents.Text(fightMap(f))),
							html.Td(
								gomponents.Text(
									fmt.Sprintf("%d", f.AlliedCount),
								),
							),
							html.Td(gomponents.Text(fightEnemies(f))),
						)
					},
				),
			),
		),
	)

	if !filtered {
		nodes = append(
			nodes,
			paginationControls(offset, total, startValue, endValue),
		)
	}

	return gomponents.Group(nodes)
}
