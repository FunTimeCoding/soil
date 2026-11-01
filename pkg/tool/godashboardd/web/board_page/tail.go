package board_page

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Tail() gomponents.Node {
	packed := Pack(p.board.Tail.Sections, p.board.Tail.Columns)
	var columns []gomponents.Node

	for _, c := range packed {
		var sections []gomponents.Node

		for _, v := range c {
			sections = append(sections, p.section(v))
		}

		columns = append(
			columns,
			html.Div(html.Class("board-column"), gomponents.Group(sections)),
		)
	}

	return grid(p.board.Tail.Columns, "board-grid board-tail", columns)
}
