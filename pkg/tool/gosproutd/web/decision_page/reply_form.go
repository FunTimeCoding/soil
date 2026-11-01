package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func replyForm(
	d *decision.Decision,
	alert string,
) gomponents.Node {
	return entry.New(actionPath(constant.ReplyPath, d.Identifier, ""), "Reply").
		Swap(target(d.Identifier), webConstant.SwapOuter).
		Inline().
		Alert(alert).
		Text(constant.ContentParameter, "Reply", "").
		Placeholder("reply without answering").
		Node()
}
