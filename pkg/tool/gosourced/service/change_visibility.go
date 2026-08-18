package service

import "github.com/funtimecoding/soil/pkg/lint/output"

func (s *Service) ChangeVisibility(
	directory string,
	symbol string,
	packagePath string,
	receiver string,
	dryRun bool,
) (*output.Results, error) {
	return s.Rename(
		directory,
		packagePath,
		symbol,
		FlipName(symbol),
		receiver,
		dryRun,
	)
}
