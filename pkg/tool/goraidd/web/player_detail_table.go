package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func playerDetailTable(rows []store.PlayerRaidRow) gomponents.Node {
	if len(rows) == 0 {
		return html.P(html.Em(gomponents.Text("No stats for this player.")))
	}

	grouped := groupByRaid(rows)
	var nodes []gomponents.Node

	for _, group := range grouped {
		first := group[0]
		nodes = append(
			nodes,
			html.H3(
				html.A(
					html.Href(
						fmt.Sprintf("/raids/%d", first.RaidID),
					),
					gomponents.Text(first.RaidName),
				),
				gomponents.Textf(" - %d fights", first.RaidFights),
			),
			playerRaidTable(group),
		)
	}

	return gomponents.Group(nodes)
}
