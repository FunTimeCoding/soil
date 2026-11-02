package repository

func (r *Repository) Declares(
	directory string,
	name string,
) bool {
	names, found := r.declarations[directory]

	if !found {
		names = declarationNames(r.Absolute(directory))
		r.declarations[directory] = names
	}

	return names[name]
}
