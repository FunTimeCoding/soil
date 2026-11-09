package repository

func (r *Repository) Tests(name string) bool {
	return len(r.testDirectories(name)) > 0
}
