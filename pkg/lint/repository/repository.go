package repository

import (
	"github.com/funtimecoding/soil/pkg/lint/repository/declaration_set"
	"github.com/funtimecoding/soil/pkg/lint/types/sibling_walk"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

type Repository struct {
	Root           string
	Files          *virtual_file_system.System
	Modules        []string
	Siblings       []string
	ImplicitBases  []string
	NestedModules  []string
	Directories    map[string]bool
	trees          *warm_forest.Forest
	contents       map[string][]string
	siblingWalks   map[string]*sibling_walk.Walk
	siblingContent map[string]string
	routes         map[string][]string
	headings       map[string][]*heading.Heading
	missing        map[string]bool
	declarations   map[string]*declaration_set.Set
	tests          map[string]map[string]bool
	ignored        func(string) bool
	siblingIgnored map[string]func(string) bool
}
