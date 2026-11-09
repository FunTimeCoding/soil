package repository

func (r *Repository) Declares(
	directory string,
	name string,
) bool {
	return r.directoryDeclarations(directory).Names[name]
}
