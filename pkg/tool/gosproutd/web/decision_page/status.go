package decision_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Status(d *decision.Decision) gomponents.Node {
	lines := []gomponents.Node{
		html.P(html.Strong(gomponents.Text(standing(d)))),
	}

	if d.Answer != "" {
		attribution := "you: "

		if d.AnswerChannel == constant.AnswerChannelConversation {
			attribution = "you, in chat, as I heard it: "
		}

		lines = append(
			lines,
			html.P(
				html.Class(webConstant.CardMetaClass),
				gomponents.Text(join.Empty(attribution, d.Answer)),
			),
		)
	}

	if d.ClearLine != "" {
		lines = append(
			lines,
			html.P(
				html.Class(webConstant.CardMetaClass),
				gomponents.Text(join.Empty("✓ taken as: ", d.ClearLine)),
			),
		)
	}

	return gomponents.Group(lines)
}
