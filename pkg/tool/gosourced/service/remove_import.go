package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) RemoveImport(
	directory string,
	filePath string,
	importPath string,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return removeImportThrough(d, filePath, importPath, out)
		},
	)
}
