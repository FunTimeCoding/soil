package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
)

func (s *Service) transact(
	x context.Context,
	directory string,
	packagePaths []string,
	dryRun bool,
	operation func(string, *sink.Sink) (*output.Results, error),
) (*output.Results, error) {
	if dryRun {
		r, e := s.attempt(directory, packagePaths, dryRun, operation)

		return r, markSettled(directory, e)
	}

	g, e := s.turns.Take(x, turn.SessionOf(x))

	if e != nil {
		return nil, e
	}

	defer g.Release()
	var moved []string
	var unloaded []string

	for range constant.CommitAttempts - 1 {
		r, f := s.attempt(directory, packagePaths, dryRun, operation)

		if failure, okay := load_failure.As(f); okay && len(unloaded) == 0 {
			unloaded = append(unloaded, failure.Package)

			continue
		}

		if f != nil {
			return nil, markSettled(directory, f)
		}

		changed := concurrentWrites(r)

		if len(changed) == 0 {
			reportTurn(r, g, moved, unloaded)

			return r, nil
		}

		moved = append(moved, changed...)
	}

	r, f := s.attempt(directory, packagePaths, dryRun, operation)

	if f != nil {
		return nil, markSettled(directory, f)
	}

	reportTurn(r, g, moved, unloaded)

	return r, nil
}
