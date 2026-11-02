package repository

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func New(
	root string,
	files *virtual_file_system.System,
) *Repository {
	result := &Repository{
		Root:           system.AbsolutePath(root),
		Files:          files,
		contents:       map[string][]string{},
		routes:         map[string][]string{},
		headings:       map[string][]*heading.Heading{},
		missing:        map[string]bool{},
		declarations:   map[string]map[string]bool{},
		siblingIgnored: map[string]func(string) bool{},
	}
	result.ignored = git.IgnoreMatcher(result.Root)
	result.Modules = modulePaths(result.Root)
	result.Siblings = siblingCheckouts(result.Root, result.Modules)

	for _, sibling := range result.Siblings {
		result.Modules = append(
			result.Modules,
			modulePaths(result.Absolute(sibling))...,
		)
		result.siblingIgnored[sibling] = git.IgnoreMatcher(
			result.Absolute(sibling),
		)
	}

	result.ImplicitBases = implicitBases(result)

	return result
}
