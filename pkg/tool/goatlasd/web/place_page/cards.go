package place_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/card"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Cards(places []*result.Place) gomponents.Node {
	cards := make([]gomponents.Node, 0, len(places))

	for _, p := range places {
		cards = append(
			cards,
			card.New(p.Name).
				Link(component.PlaceLocator(p.Kind, p.Name)).
				Meta(
					fmt.Sprintf(
						constant.CardKindFormat,
						component.Word(
							p.Count,
							constant.PlacementWord,
							constant.PlacementsWord,
						),
						component.PlaceLabel(p.Kind),
					),
				).
				State(
					html.P(
						html.Class(constant.PlaceCountClass),
						gomponents.Textf("%d", p.Count),
					),
				).
				Terms(p.Name).
				Node(),
		)
	}

	return tile_grid.New(cards).Node()
}
