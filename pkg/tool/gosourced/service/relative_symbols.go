package service

import (
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"path/filepath"
)

func relativeSymbols(
	root string,
	moves []*symbol_move.Move,
) []*symbol_move.Move {
	var result []*symbol_move.Move

	for _, m := range moves {
		directory, e := filepath.Rel(root, m.Directory)

		if e != nil {
			continue
		}

		newDirectory, f := filepath.Rel(root, m.NewDirectory)

		if f != nil {
			continue
		}

		result = append(
			result,
			symbol_move.New(
				filepath.ToSlash(directory),
				m.Name,
				filepath.ToSlash(newDirectory),
				m.NewName,
			),
		)
	}

	return result
}
