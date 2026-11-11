package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/roster_card"
	"github.com/funtimecoding/soil/pkg/web/component/card"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func rosterCard(c *roster_card.Card) gomponents.Node {
	identifier := c.Session.Identifier

	return card.New(c.Session.Name).
		Identifier(Mark(identifier)).
		Preserve().
		Link(Path(identifier)).
		Marker(
			html.Span(
				extended.StreamSwap(
					Event(identifier, constant.RosterActivityPart),
				),
				ActivityDot(c),
			),
		).
		State(region(identifier, constant.RosterStatePart, State(c))).
		Footer(RosterCruise(identifier, c.Cruise, c.CruiseGrace, "")).
		Node()
}
