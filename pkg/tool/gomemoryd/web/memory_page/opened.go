package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Opened() gomponents.Node {
	m := p.view.Memory
	r := p.row()
	acts := []gomponents.Node{
		load.New(actionPath(m.Identifier, constant.CloseAction), "Close").
			Swap(r.Target(), webConstant.SwapOuter).
			Button().
			Class("outline secondary").
			Node(),
		load.New(Path(m.Identifier), "Open page").
			Button().
			Class("outline secondary").
			Node(),
	}

	if m.ProvenanceFile == "" {
		acts = append(
			[]gomponents.Node{
				load.New(actionPath(m.Identifier, constant.EditAction), "Edit").
					Swap(r.Target(), webConstant.SwapOuter).
					Button().
					Class("outline").
					Node(),
			},
			acts...,
		)
	}

	return r.Detail(
		len(cells(m.Name, m.Description, m.Tags)),
		p.facts(),
		gomponents.If(
			m.ProvenanceFile != "",
			html.P(
				html.Small(
					html.Class(webConstant.AbsentClass),
					gomponents.Textf(
						"Read-only, loaded from %s; edit it there.",
						m.ProvenanceFile,
					),
				),
			),
		),
		html.Pre(gomponents.Text(m.Content)),
		html.Div(html.Class(webConstant.ActionRowClass), gomponents.Group(acts)),
	).Response()
}
