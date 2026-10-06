package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"slices"
)

func (s *Service) commit(
	directory string,
	before *snapshot.Snapshot,
	roots []string,
	outs []*sink.Sink,
	r *output.Results,
) (*output.Results, error) {
	if !slices.ContainsFunc(
		outs,
		func(o *sink.Sink) bool {
			return !o.Empty()
		},
	) {
		return r, nil
	}
	s.commits.Lock()
	defer s.commits.Unlock()

	if s.beforeCommit != nil {
		s.beforeCommit()
	}

	if changed := before.Changed(roots...); len(changed) > 0 {
		refused := output.NewResultsWithDirectory(directory)

		for _, path := range changed {
			refused.AddConcern(
				concern.NewFile(
					constant.ConcernConcurrentWrite,
					"changed since it was read - nothing written, run again",
					path,
					false,
				),
			)
		}

		return refused, nil
	}

	for _, out := range outs {
		if e := out.Commit(); e != nil {
			return nil, e
		}
	}

	return r, nil
}
