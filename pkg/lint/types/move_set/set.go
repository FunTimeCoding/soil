package move_set

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
)

type Set struct {
	Paths   []*path_move.Move
	Symbols []*symbol_move.Move
}
