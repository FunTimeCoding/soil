package repository

func (r *Repository) DeclaresType(
	directory string,
	name string,
) bool {
	return r.directoryDeclarations(directory).Types[name]
}
