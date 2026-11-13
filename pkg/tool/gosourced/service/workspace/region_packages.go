package workspace

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"golang.org/x/tools/go/packages"
	"strings"
)

func (s *Service) RegionPackages(
	directory string,
	region string,
) ([]*packages.Package, error) {
	if s.full {
		all, _, e := load.Packages(directory, goModConstant.AllPackages)

		return all, e
	}

	var patterns []string

	for path, u := range s.Workspace(directory).Graph().Units {
		if strings.HasPrefix(path, region) {
			patterns = append(patterns, u.Directory)
		}
	}

	if len(patterns) == 0 {
		return nil, nil
	}

	all, _, e := load.Packages(directory, patterns...)

	return all, e
}
