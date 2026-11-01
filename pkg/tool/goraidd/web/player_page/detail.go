package player_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(rows []record.PlayerRaid) gomponents.Node {
	if len(rows) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("No stats for this player."),
		)
	}

	var nodes []gomponents.Node

	for _, g := range group(rows) {
		first := g[0]
		nodes = append(
			nodes,
			html.H3(
				html.A(
					html.Href(component.RaidPath(first.RaidIdentifier)),
					gomponents.Text(first.RaidName),
				),
				gomponents.Textf(" - %d fights", first.RaidFights),
			),
			raidTable(g),
		)
	}

	return gomponents.Group(nodes)
}
