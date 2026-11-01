package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Detail() gomponents.Node {
	m := p.view.Memory

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(m.Name)),
			html.Div(
				html.Class(constant.DetailHeadClass),
				p.facts(),
				gomponents.If(
					m.Description != "",
					html.P(gomponents.Text(m.Description)),
				),
			),
			TagChips(m.Tags),
			html.H4(gomponents.Text("Content")),
			html.Pre(gomponents.Text(m.Content)),
			p.history(),
			p.related(),
		},
	)
}
