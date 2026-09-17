package repository

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func New(
	root string,
	files *virtual_file_system.System,
) *Repository {
	result := &Repository{
		Root:     system.AbsolutePath(root),
		Files:    files,
		contents: map[string][]string{},
		routes:   map[string][]string{},
		missing:  map[string]bool{},
	}
	result.Modules = modulePaths(result.Root)
	result.Siblings = siblingCheckouts(result.Root, result.Modules)

	for _, sibling := range result.Siblings {
		result.Modules = append(
			result.Modules,
			modulePaths(result.Absolute(sibling))...,
		)
	}

	result.ImplicitBases = implicitBases(result)

	return result
}
