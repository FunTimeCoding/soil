package dashboard_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Tags() gomponents.Node {
	heading := html.H3(gomponents.Text("Tags"))

	if len(p.view.Tags) == 0 {
		return gomponents.Group(
			[]gomponents.Node{
				heading,
				html.P(
					html.Class(webConstant.AbsentClass),
					gomponents.Text("No tags."),
				),
			},
		)
	}

	chips := make([]gomponents.Node, 0, len(p.view.Tags))

	for _, t := range p.view.Tags {
		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s (%d)", t.Tag, t.Count)).
				Link(memory_page.TagPath(t.Tag)).
				Node(),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			heading,
			html.Div(
				html.Class(webConstant.ClusterClass),
				gomponents.Group(chips),
			),
		},
	)
}
