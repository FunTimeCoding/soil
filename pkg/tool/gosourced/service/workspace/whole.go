package workspace

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

func (s *Service) Whole(
	directory string,
	packagePath string,
) ([]*packages.Package, *token.FileSet, error) {
	return wholeLoad(directory, s.Workspace(directory), packagePath)
}
