package decision_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/card"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Hero(
	d *decision.Decision,
	answerAlert string,
	replyAlert string,
) gomponents.Node {
	return meta(card.New(d.Question).Task().Hero(), d).
		Meta(join.Empty("If you never answer: ", d.DefaultAction)).
		State(
			region(d.Identifier, constant.StatusPart, Status(d)),
			answerForm(d, answerAlert),
			region(d.Identifier, constant.ThreadPart, Thread(d)),
			replyForm(d, replyAlert),
		).
		Footer(
			buttonRegion(d.Identifier, constant.ChoicePart, Choices(d)),
			buttonRegion(d.Identifier, constant.DismissPart, Dismissals(d)),
			load.New(Path(d.Identifier, false), "Close").
				Swap(target(d.Identifier), webConstant.SwapOuter).
				Button().
				Class("outline secondary").
				Node(),
		).
		Node()
}
