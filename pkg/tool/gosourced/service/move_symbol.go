package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func (s *Service) MoveSymbol(
	x context.Context,
	directory string,
	packagePath string,
	symbol string,
	targetPackagePath string,
	targetFile string,
	create bool,
	dryRun bool,
) (*output.Results, error) {
	return s.MoveSymbols(
		x,
		directory,
		packagePath,
		[]string{symbol},
		"",
		targetPackagePath,
		targetFile,
		create,
		false,
		dryRun,
	)
}
