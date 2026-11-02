package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
)

func symbolMoves(moves ...*symbol_move.Move) *move_set.Set {
	return move_set.New(nil, moves)
}
