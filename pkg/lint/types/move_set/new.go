package move_set

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
)

func New(
	paths []*path_move.Move,
	symbols []*symbol_move.Move,
) *Set {
	return &Set{Paths: paths, Symbols: symbols}
}
