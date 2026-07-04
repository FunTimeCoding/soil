package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) topGrid(values map[string][]string) gomponents.Node {
	var columns []gomponents.Node

	for _, c := range s.board.Top {
		var sections []gomponents.Node

		for _, v := range c.Sections {
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

	return boardGrid(len(s.board.Top), "board-grid", columns)
}
