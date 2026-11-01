package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Grid(decisions []*decision.Decision) gomponents.Node {
	if len(decisions) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("This session has pushed nothing."),
		)
	}

	cards := make([]gomponents.Node, 0, len(decisions))

	for _, d := range decisions {
		cards = append(cards, decision_page.Preserved(d))
	}

	return tile_grid.New(cards).Node()
}
