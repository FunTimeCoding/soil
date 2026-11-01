package decision_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Thread(d *decision.Decision) gomponents.Node {
	turns := make([]gomponents.Node, 0, len(d.Turns))

	for _, v := range d.Turns {
		turns = append(
			turns,
			html.P(
				html.Small(
					html.Strong(
						gomponents.Text(join.Empty(string(v.Author), ": ")),
					),
					gomponents.Text(v.Content),
				),
			),
		)
	}

	return gomponents.Group(turns)
}
