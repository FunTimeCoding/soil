package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) RemoveParameters(
	directory string,
	removals []*removal.Parameter,
	dryRun bool,
) (*output.Results, error) {
	var paths []string

	for _, p := range removals {
		paths = append(paths, p.PackagePath)
	}

	return s.transact(
		directory,
		paths,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return s.removeParametersThrough(d, removals, out)
		},
	)
}
