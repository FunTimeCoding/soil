package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Choices(d *decision.Decision) gomponents.Node {
	if d.ClearLine != "" {
		return gomponents.Group(nil)
	}

	buttons := make([]gomponents.Node, 0, len(d.Choices))

	for _, v := range d.Choices {
		buttons = append(
			buttons,
			action.New(
				actionPath(
					constant.AnswerPath,
					d.Identifier,
					constant.AnswerKindChoice,
				),
				v.Label,
			).
				Hidden(constant.ValueParameter, v.Label).
				Swap(target(d.Identifier), webConstant.SwapOuter).
				Node(),
		)
	}

	return gomponents.Group(buttons)
}
