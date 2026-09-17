package repository

import "github.com/funtimecoding/soil/pkg/system"

func (r *Repository) SiblingExists(path string) bool {
	full := r.Absolute(path)

	return system.FileExists(full) || system.DirectoryExists(full)
}
