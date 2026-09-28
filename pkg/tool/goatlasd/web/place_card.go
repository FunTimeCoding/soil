package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func placeCard(p *result.Place) gomponents.Node {
	return html.Div(
		html.Class(constant.PlaceCardClass),
		html.A(
			html.Href(placeLocator(p.Kind, p.Name)),
			html.H4(gomponents.Text(p.Name)),
			html.P(
				html.Class(constant.PlaceCountClass),
				gomponents.Textf("%d", p.Count),
			),
			html.P(
				html.Class(constant.PlaceKindClass),
				gomponents.Textf(
					constant.CardKindFormat,
					Word(
						p.Count,
						constant.PlacementWord,
						constant.PlacementsWord,
					),
					placeLabel(p.Kind),
				),
			),
		),
	)
}
