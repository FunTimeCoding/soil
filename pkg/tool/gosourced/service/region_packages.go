package service

import (
	"golang.org/x/tools/go/packages"
	"strings"
)

func (s *Service) regionPackages(
	directory string,
	region string,
) ([]*packages.Package, error) {
	if s.full {
		all, _, e := loadPackages(directory, "./...")

		return all, e
	}

	var patterns []string

	for path, u := range s.workspace(directory).Graph().Units {
		if strings.HasPrefix(path, region) {
			patterns = append(patterns, u.Directory)
		}
	}

	if len(patterns) == 0 {
		return nil, nil
	}

	all, _, e := loadPackages(directory, patterns...)

	return all, e
}
