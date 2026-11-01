package counter_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Note() gomponents.Node {
	return html.P(
		html.Class(constant.CounterNoteClass),
		html.Small(
			gomponents.Text(
				"Irrelevant counts my bad asks. Abandoned counts the ones you never needed to answer. Both are measurements of me, not of you.",
			),
		),
	)
}
