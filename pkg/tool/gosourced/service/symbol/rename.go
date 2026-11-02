package symbol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
)

func (s *Service) Rename(
	x context.Context,
	directory string,
	packagePath string,
	oldName string,
	newName string,
	receiver string,
	dryRun bool,
) (*output.Results, error) {
	return s.RenameSymbols(
		x,
		directory,
		packagePath,
		[]*symbol_rename.Rename{symbol_rename.New(oldName, newName, receiver)},
		dryRun,
	)
}
