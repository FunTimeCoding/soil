package board_page

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

type Page struct {
	board  *board.Board
	values map[string][]string
}
