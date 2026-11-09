package repository

import "github.com/funtimecoding/soil/pkg/lint/repository/declaration_set"

func (r *Repository) directoryDeclarations(directory string) *declaration_set.Set {
	result, found := r.declarations[directory]

	if !found {
		result = parseDeclarations(r.Absolute(directory))
		r.declarations[directory] = result
	}

	return result
}
