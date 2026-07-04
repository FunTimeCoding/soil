package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) tailGrid(values map[string][]string) gomponents.Node {
	packed := pack(s.board.Tail.Sections, s.board.Tail.Columns)
	var columns []gomponents.Node

	for _, c := range packed {
		var sections []gomponents.Node

		for _, v := range c {
			sections = append(sections, sectionCard(v, values))
		}

		columns = append(
			columns,
			html.Div(
				html.Class("board-column"),
				gomponents.Group(sections),
			),
		)
	}

	return boardGrid(
		s.board.Tail.Columns,
		"board-grid board-tail",
		columns,
	)
}
