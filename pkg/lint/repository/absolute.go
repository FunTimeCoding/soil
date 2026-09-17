package repository

import "path/filepath"

func (r *Repository) Absolute(path string) string {
	return filepath.Join(r.Root, path)
}
