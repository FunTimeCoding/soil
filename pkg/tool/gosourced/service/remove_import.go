package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) RemoveImport(
	x context.Context,
	directory string,
	filePath string,
	importPath string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		x,
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return removeImportThrough(d, filePath, importPath, out)
		},
	)
}
