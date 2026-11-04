package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
)

func (s *Service) AddParameters(
	x context.Context,
	directory string,
	additions []*addition.Parameter,
	imports []string,
	dryRun bool,
) (*output.Results, error) {
	var paths []string

	for _, a := range additions {
		paths = append(paths, a.PackagePath)
	}

	return s.commit.Transact(
		x,
		directory,
		paths,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return s.addParametersThrough(d, additions, imports, out)
		},
	)
}
