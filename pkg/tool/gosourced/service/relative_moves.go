package service

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
	"path/filepath"
)

func relativeMoves(
	root string,
	renames []*sink_operation.Operation,
) []*path_move.Move {
	var result []*path_move.Move

	for _, o := range renames {
		from, e := filepath.Rel(root, o.Path)

		if e != nil {
			continue
		}

		to, f := filepath.Rel(root, o.Target)

		if f != nil {
			continue
		}

		result = append(
			result,
			path_move.New(filepath.ToSlash(from), filepath.ToSlash(to)),
		)
	}

	return result
}
