package repository

import "github.com/funtimecoding/soil/pkg/system"

func (r *Repository) Exists(path string) bool {
	if path == "" {
		return false
	}

	if r.Files.Has(path) ||
		r.Files.DirectoryExists(path) ||
		system.DirectoryExists(r.Absolute(path)) {
		return true
	}

	return r.inNestedModule(path) && system.FileExists(r.Absolute(path))
}
