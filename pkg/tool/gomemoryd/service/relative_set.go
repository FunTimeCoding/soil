package service

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
)

func relativeSet(
	root string,
	moves *move_set.Set,
) *move_set.Set {
	var paths []*path_move.Move

	for _, m := range moves.Paths {
		paths = append(
			paths,
			path_move.New(rootRelative(root, m.From), rootRelative(root, m.To)),
		)
	}

	var symbols []*symbol_move.Move

	for _, m := range moves.Symbols {
		symbols = append(
			symbols,
			symbol_move.New(
				rootRelative(root, m.Directory),
				m.Name,
				rootRelative(root, m.NewDirectory),
				m.NewName,
			),
		)
	}

	return move_set.New(paths, symbols)
}
