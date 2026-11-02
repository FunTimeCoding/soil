package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
)

func pathMoves(moves ...*path_move.Move) *move_set.Set {
	return move_set.New(moves, nil)
}
