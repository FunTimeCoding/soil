package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/card"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func compact(d *decision.Decision) *card.Card {
	return meta(card.New(d.Question).Task(), d).
		State(region(d.Identifier, constant.StatusPart, Status(d))).
		Footer(
			buttonRegion(d.Identifier, constant.ChoicePart, Choices(d)),
			load.New(Path(d.Identifier, true), "Open…").
				Swap(target(d.Identifier), webConstant.SwapOuter).
				Button().
				Class("outline").
				Node(),
		)
}
