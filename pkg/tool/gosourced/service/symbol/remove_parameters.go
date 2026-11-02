package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"
)

func (s *Service) RemoveParameters(
	x context.Context,
	directory string,
	removals []*removal.Parameter,
	dryRun bool,
) (*output.Results, error) {
	var paths []string

	for _, p := range removals {
		paths = append(paths, p.PackagePath)
	}

	return s.commit.Transact(
		x,
		directory,
		paths,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return s.removeParametersThrough(d, removals, out)
		},
	)
}
