package service

import "github.com/funtimecoding/soil/pkg/lint/output"

func (s *Service) MoveSymbol(
	directory string,
	packagePath string,
	symbol string,
	targetPackagePath string,
	targetFile string,
	create bool,
) (*output.Results, error) {
	return s.MoveSymbols(
		directory,
		packagePath,
		[]string{symbol},
		"",
		targetPackagePath,
		targetFile,
		create,
	)
}
