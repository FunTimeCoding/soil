package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) RenameModule(
	x context.Context,
	directory string,
	modulePath string,
	newModulePath string,
	force bool,
	dryRun bool,
) (*output.Results, error) {
	return s.transact(
		x,
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return renameModuleThrough(
				d,
				modulePath,
				newModulePath,
				force,
				dryRun,
				out,
			)
		},
	)
}
