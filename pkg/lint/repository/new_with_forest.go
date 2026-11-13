package repository

import (
	"github.com/funtimecoding/soil/pkg/lint/repository/declaration_set"
	"github.com/funtimecoding/soil/pkg/lint/types/sibling_walk"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func NewWithForest(
	root string,
	files *virtual_file_system.System,
	trees *warm_forest.Forest,
) *Repository {
	result := &Repository{
		Root:           system.AbsolutePath(root),
		Files:          files,
		trees:          trees,
		contents:       map[string][]string{},
		routes:         map[string][]string{},
		headings:       map[string][]*heading.Heading{},
		missing:        map[string]bool{},
		declarations:   map[string]*declaration_set.Set{},
		siblingIgnored: map[string]func(string) bool{},
		siblingWalks:   map[string]*sibling_walk.Walk{},
		siblingContent: map[string]string{},
	}
	result.ignored = trees.Tree(result.Root).IgnoreMatcher()
	result.Modules = modulePaths(result.Root)
	result.Siblings = siblingCheckouts(result.Root, result.Modules)

	for _, sibling := range result.Siblings {
		result.Modules = append(
			result.Modules,
			modulePaths(result.Absolute(sibling))...,
		)
		result.siblingIgnored[sibling] = trees.Tree(
			result.Absolute(sibling),
		).IgnoreMatcher()
	}

	result.ImplicitBases = implicitBases(result)

	return result
}
