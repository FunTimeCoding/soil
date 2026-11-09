package reference

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) RetargetReferences(
	x context.Context,
	directory string,
	moves []*path_move.Move,
	dryRun bool,
) (*output.Results, error) {
	refused := output.NewResultsWithDirectory(directory)

	if len(moves) == 0 {
		return commit.FailValidation(refused, "no moves given")
	}

	for _, m := range moves {
		for _, p := range []string{m.Old, m.New} {
			if !system.InsideDirectory(directory, modulePath(directory, p)) {
				return commit.FailValidation(
					refused,
					fmt.Sprintf("%s is outside the module", p),
				)
			}
		}
	}

	r, e := s.commit.Transact(
		x,
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			for _, m := range moves {
				out.Relocate(modulePath(d, m.Old), modulePath(d, m.New))
			}

			return output.NewResultsWithDirectory(d), nil
		},
	)

	if e != nil || len(r.Entries) > 0 {
		return r, e
	}

	r.AddConcern(
		concern.NewFile(
			constant.ConcernRetargeted,
			"nothing references the moved paths",
			"",
			true,
		),
	)

	return r, nil
}
