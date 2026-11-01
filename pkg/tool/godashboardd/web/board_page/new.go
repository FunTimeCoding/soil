package board_page

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

func New(
	b *board.Board,
	values map[string][]string,
) *Page {
	return &Page{board: b, values: values}
}
