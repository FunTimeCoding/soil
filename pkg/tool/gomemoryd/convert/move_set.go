package convert

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func MoveSet(r *server.RetargetRequest) *move_set.Set {
	var paths []*path_move.Move

	for _, m := range r.Paths {
		paths = append(paths, path_move.New(m.From, m.To))
	}

	var symbols []*symbol_move.Move

	for _, m := range r.Symbols {
		symbols = append(
			symbols,
			symbol_move.New(m.Directory, m.Name, m.NewDirectory, m.NewName),
		)
	}

	return move_set.New(paths, symbols)
}
