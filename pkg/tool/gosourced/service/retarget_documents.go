package service

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/sink_operation"
	"path/filepath"
	"slices"
)

func (s *Service) retargetDocuments(
	directory string,
	outs []*sink.Sink,
	r *output.Results,
) []*sink.Sink {
	var renames []*sink_operation.Operation
	var symbols []*symbol_move.Move

	for _, o := range outs {
		renames = append(renames, o.Renames()...)
		symbols = append(symbols, o.Symbols()...)
	}

	if len(renames) == 0 && len(symbols) == 0 {
		return outs
	}

	result := outs

	for _, root := range slices.Concat(
		[]string{directory},
		s.inventory.Requiring(directory),
	) {
		documents := lint.RetargetTree(
			root,
			move_set.New(
				relativeMoves(root, renames),
				relativeSymbols(root, symbols),
			),
		)

		if len(documents) == 0 {
			continue
		}

		out := sinkFor(result, root)

		if out == nil {
			out = sink.New(root)
			result = append(result, out)
		}

		for _, d := range documents {
			path := filepath.Join(root, d.Path)

			if d.Content != d.Original {
				out.Rewrite(path, []byte(d.Original), []byte(d.Content))
			}

			for _, c := range d.Changes {
				r.AddConcern(retargetConcern(path, c))
			}
		}
	}

	return result
}
