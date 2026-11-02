package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
)

func (s *Service) RenameSymbols(
	x context.Context,
	directory string,
	packagePath string,
	renames []*symbol_rename.Rename,
	dryRun bool,
) (*output.Results, error) {
	refused := output.NewResultsWithDirectory(directory)

	if len(renames) == 0 {
		return failValidation(refused, "no renames given")
	}

	if message := duplicateTargets(renames); message != "" {
		return failValidation(refused, message)
	}

	return s.transact(
		x,
		directory,
		[]string{packagePath},
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			return s.renameThrough(d, packagePath, renames, out)
		},
	)
}
