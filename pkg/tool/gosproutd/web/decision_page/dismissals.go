package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Dismissals(d *decision.Decision) gomponents.Node {
	if d.ClearLine != "" {
		return gomponents.Group(nil)
	}

	var buttons []gomponents.Node

	for _, v := range []constant.State{
		constant.StateDeclined,
		constant.StateIrrelevant,
		constant.StatePostponed,
	} {
		buttons = append(
			buttons,
			action.New(
				actionPath(constant.DismissPath, d.Identifier, ""),
				string(v),
			).
				Hidden(constant.StateParameter, string(v)).
				Swap(target(d.Identifier), webConstant.SwapOuter).
				ButtonClass("outline secondary").
				Node(),
		)
	}

	return gomponents.Group(buttons)
}
