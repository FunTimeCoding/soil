package dashboard_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/relation_page"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Memories() gomponents.Node {
	var rows []gomponents.Node

	for _, t := range constant.MemoryTypes {
		if c, okay := p.view.Types[t]; okay {
			rows = append(
				rows,
				html.Tr(
					html.Td(gomponents.Text(t)),
					html.Td(
						html.Class(webConstant.NumericClass),
						gomponents.Textf("%d", c),
					),
				),
			)
		}
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.MemoriesTitle)),
			table.New("Type", "Count").
				Numeric(1).
				Rows(rows).
				Empty("No memories.").
				Node(),
			gomponents.If(
				p.view.Untyped > 0,
				html.Div(
					html.Class(webConstant.ClusterClass),
					chip.New(
						fmt.Sprintf("%d untyped relations", p.view.Untyped),
					).
						Link(relation_page.Path(constant.UntypedFilter)).
						Node(),
				),
			),
		},
	)
}
