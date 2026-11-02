package module

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) RenameModule(
	x context.Context,
	directory string,
	modulePath string,
	newModulePath string,
	force bool,
	dryRun bool,
) (*output.Results, error) {
	return s.commit.Transact(
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
