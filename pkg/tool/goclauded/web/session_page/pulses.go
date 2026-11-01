package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Pulses(pulses []pulse.Pulse) gomponents.Node {
	if len(pulses) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("No pulses."),
		)
	}

	lines := make([]gomponents.Node, 0, len(pulses))

	for _, p := range pulses {
		from := "→"

		if p.FromName != "" {
			from = fmt.Sprintf("%s →", p.FromName)
		}

		lines = append(
			lines,
			html.P(
				html.Small(gomponents.Textf("%s ", p.CreatedAt.Format("15:04"))),
				html.Em(gomponents.Textf("%s ", from)),
				gomponents.Text(p.Body),
			),
		)
	}

	return gomponents.Group(lines)
}
