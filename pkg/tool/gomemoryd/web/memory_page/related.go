package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/component"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) related() gomponents.Node {
	if len(p.view.Related) == 0 {
		return gomponents.Group(nil)
	}

	scope := p.view.Memory.Scope
	items := make([]gomponents.Node, 0, len(p.view.Related))

	for _, r := range p.view.Related {
		items = append(
			items,
			html.Li(
				html.A(html.Href(Path(r.Identifier)), gomponents.Text(r.Name)),
				gomponents.If(
					r.Type != "",
					html.Small(gomponents.Textf(" (%s)", r.Type)),
				),
				gomponents.If(
					r.Scope != scope,
					html.Small(
						gomponents.Textf(" [%s]", component.ScopeName(r.Scope)),
					),
				),
				gomponents.Textf(" - %s", r.Description),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H4(gomponents.Text(constant.RelationsTitle)),
			html.Ul(items...),
		},
	)
}
