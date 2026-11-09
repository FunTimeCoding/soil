package repository

import "path/filepath"

func (r *Repository) TestElsewhere(
	name string,
	directory string,
) bool {
	moved := filepath.Clean(r.Absolute(directory))

	for d := range r.testDirectories(name) {
		if d != moved {
			return true
		}
	}

	return false
}
