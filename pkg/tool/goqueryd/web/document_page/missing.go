package document_page

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Missing(
	path string,
	suggestions []string,
) gomponents.Node {
	items := make([]gomponents.Node, 0, len(suggestions))

	for _, s := range suggestions {
		items = append(
			items,
			html.Li(html.A(html.Href(Path(s)), gomponents.Text(s))),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Not Found")),
			html.P(
				html.Class(webConstant.AbsentClass),
				gomponents.Textf("No document at: %s", path),
			),
			gomponents.If(
				len(items) > 0,
				gomponents.Group(
					[]gomponents.Node{
						html.P(gomponents.Text("Did you mean:")),
						html.Ul(items...),
					},
				),
			),
		},
	)
}
