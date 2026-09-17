package repository

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

type Repository struct {
	Root          string
	Files         *virtual_file_system.System
	Modules       []string
	Siblings      []string
	ImplicitBases []string
	contents      map[string][]string
	routes        map[string][]string
	missing       map[string]bool
}
