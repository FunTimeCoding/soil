package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Roster(
	sessions []session.Session,
	labels map[string][]label.Label,
	pulses map[string]*pulse.Pulse,
) gomponents.Node {
	if len(sessions) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("No active sessions."),
		)
	}

	cards := make([]gomponents.Node, 0, len(sessions))

	for i := range sessions {
		s := &sessions[i]
		cards = append(
			cards,
			rosterCard(s, labels[s.Identifier], pulses[s.Identifier]),
		)
	}

	return tile_grid.New(cards).Node()
}
