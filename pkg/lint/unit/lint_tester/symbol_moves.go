package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
)

func SymbolMoves(moves ...*symbol_move.Move) *move_set.Set {
	return move_set.New(nil, moves)
}
