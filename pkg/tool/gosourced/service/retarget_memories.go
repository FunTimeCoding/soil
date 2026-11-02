package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) retargetMemories(
	outs []*sink.Sink,
	dryRun bool,
	r *output.Results,
) {
	var paths []*path_move.Move
	var symbols []*symbol_move.Move

	for _, o := range outs {
		for _, rename := range o.Renames() {
			paths = append(paths, path_move.New(rename.Path, rename.Target))
		}

		symbols = append(symbols, o.Symbols()...)
	}

	if len(paths) == 0 && len(symbols) == 0 {
		return
	}

	var reports []*retarget_report.Report
	e := s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			var f error
			reports, f = s.memory.PostReferencesRetarget(
				x,
				move_set.New(paths, symbols),
				dryRun,
			)

			return f
		},
	)

	if e != nil {
		if classify.Reportable(e) {
			s.reporter.CaptureException(e)
		}

		r.AddConcern(
			concern.NewFile(
				constant.ConcernReferenceLeft,
				fmt.Sprintf("memories not retargeted: %s", e),
				"",
				true,
			),
		)

		return
	}

	for _, report := range reports {
		for _, c := range report.Changes {
			r.AddConcern(memoryConcern(report, c))
		}
	}
}
