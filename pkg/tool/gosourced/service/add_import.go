package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) AddImport(
	x context.Context,
	directory string,
	filePath string,
	importPath string,
	alias string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		x,
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return addImportThrough(d, filePath, importPath, alias, out)
		},
	)
}
