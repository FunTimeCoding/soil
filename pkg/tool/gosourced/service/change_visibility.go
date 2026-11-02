package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func (s *Service) ChangeVisibility(
	x context.Context,
	directory string,
	symbol string,
	packagePath string,
	receiver string,
	dryRun bool,
) (*output.Results, error) {
	return s.Rename(
		x,
		directory,
		packagePath,
		symbol,
		FlipName(symbol),
		receiver,
		dryRun,
	)
}
