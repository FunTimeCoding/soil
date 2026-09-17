package repository

import "github.com/funtimecoding/soil/pkg/system"

func (r *Repository) Exists(path string) bool {
	return r.Files.Has(path) ||
		r.Files.DirectoryExists(path) ||
		system.DirectoryExists(r.Absolute(path))
}
