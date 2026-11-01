package board_page

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Top() gomponents.Node {
	var columns []gomponents.Node

	for _, c := range p.board.Top {
		var sections []gomponents.Node

		for _, v := range c.Sections {
			sections = append(sections, p.section(v))
		}

		columns = append(
			columns,
			html.Div(html.Class("board-column"), gomponents.Group(sections)),
		)
	}

	return grid(len(p.board.Top), "board-grid", columns)
}
