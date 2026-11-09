package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/roster_card"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Roster(cards []*roster_card.Card) gomponents.Node {
	if len(cards) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("No active sessions."),
		)
	}

	nodes := make([]gomponents.Node, 0, len(cards))

	for _, c := range cards {
		nodes = append(nodes, rosterCard(c))
	}

	return tile_grid.New(nodes).Node()
}
