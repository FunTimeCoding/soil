package repository

import (
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

type Repository struct {
	Root          string
	Files         *virtual_file_system.System
	Modules       []string
	Siblings      []string
	ImplicitBases []string
	contents      map[string][]string
	routes        map[string][]string
	headings      map[string][]*heading.Heading
	missing       map[string]bool
}
