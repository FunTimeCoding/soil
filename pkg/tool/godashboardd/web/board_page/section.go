package board_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"
	"github.com/funtimecoding/soil/pkg/web/component/panel"
	"maragu.dev/gomponents"
)

func (p *Page) section(v *layout.Section) gomponents.Node {
	result := panel.New(v.Name).Class("board-section")

	for _, e := range v.Entries {
		result.Content(p.entry(e))
	}

	return result.Node()
}
