package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) transact(
	directory string,
	packagePaths []string,
	dryRun bool,
	operation func(string, *sink.Sink) (*output.Results, error),
) (*output.Results, error) {
	candidates := []string{directory}

	if len(packagePaths) > 0 {
		candidates = append(candidates, s.inventory.Replacing(directory)...)
	}

	before := snapshot.Take(candidates...)
	roots := []string{directory}
	var captures []*sink.Sink
	var reached []*concern.Concern

	for _, other := range s.reaching(directory, packagePaths) {
		capture := sink.New(other)
		r, e := operation(other, capture)

		if e != nil {
			return nil, e
		}

		if hasUnfixed(r) {
			refused := output.NewResultsWithDirectory(directory)

			for _, c := range rebaseConcerns(other, r) {
				refused.AddConcern(c)
			}

			return refused, nil
		}

		for _, c := range rebaseConcerns(other, r) {
			if system.InsideDirectory(other, c.Path) {
				reached = append(reached, c)
			}
		}

		roots = append(roots, other)
		captures = append(captures, capture)
	}

	out := sink.New(directory)
	r, e := operation(directory, out)

	if e != nil || hasUnfixed(r) {
		return r, e
	}

	if dropped := out.Dropped(); len(dropped) > 0 {
		return refuseOutside(directory, dropped), nil
	}

	for _, c := range reached {
		r.AddConcern(c)
	}

	if dryRun {
		r.MarkPlanned()

		return r, nil
	}

	return s.commit(
		directory,
		before,
		roots,
		append([]*sink.Sink{out}, captures...),
		r,
	)
}
