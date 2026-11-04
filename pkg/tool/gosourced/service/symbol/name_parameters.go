package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
)

func (s *Service) NameParameters(
	x context.Context,
	directory string,
	request *naming.Request,
	dryRun bool,
) (*output.Results, error) {
	paths := s.scopePaths(directory, request.PackagePath)

	return s.commit.Transact(
		x,
		directory,
		paths,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return s.nameParametersThrough(d, request, paths, out)
		},
	)
}
