package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Modes() gomponents.Node {
	var chips []gomponents.Node

	for _, m := range []string{constant.KeywordMode, constant.HybridMode} {
		c := chip.New(m).Link(p.path(m, p.metadata))

		if m == p.mode {
			c.Current()
		}

		chips = append(chips, c.Node())
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Text("Mode:"),
		gomponents.Group(chips),
	)
}
