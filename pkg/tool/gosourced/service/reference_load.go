package service

import (
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func (s *Service) referenceLoad(
	directory string,
	packagePath string,
	targetPackagePath string,
	objects func([]*packages.Package, *token.FileSet) []types.Object,
) ([]*packages.Package, *token.FileSet, error) {
	w := s.workspace(directory)

	if s.full {
		return wholeLoad(directory, w, packagePath, targetPackagePath)
	}

	i := w.References()
	reached := reachedPaths(w.Graph(), []string{packagePath, targetPackagePath})

	if !i.Has(packagePath) && !w.Reaches(packagePath) ||
		slices.ContainsFunc(
			i.Unindexed(),
			func(unit string) bool { return reached[unit] },
		) {
		return wholeLoad(directory, w, packagePath, targetPackagePath)
	}

	patterns := []string{packagePath}

	if targetPackagePath != "" && i.Has(targetPackagePath) {
		patterns = append(patterns, targetPackagePath)
	}

	first, set, e := loadPackages(directory, patterns...)

	if e != nil {
		return nil, nil, e
	}

	targets := objects(first, set)

	if len(targets) == 0 {
		return first, set, nil
	}

	loaded := len(patterns)

	for _, o := range targets {
		t, okay := xref.Target(o)

		if !okay {
			return wholeLoad(directory, w, packagePath, targetPackagePath)
		}

		for _, unit := range i.Referencing(t) {
			patterns = append(patterns, i.Directory(unit))
		}
	}

	if len(patterns) == loaded {
		return first, set, nil
	}

	patterns = slices.Compact(slices.Sorted(slices.Values(patterns)))

	if 2*len(patterns) > i.Count() {
		return wholeLoad(directory, w, packagePath, targetPackagePath)
	}

	return loadPackages(directory, patterns...)
}
