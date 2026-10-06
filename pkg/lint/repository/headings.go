package repository

import (
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"github.com/funtimecoding/soil/pkg/system"
)

func (r *Repository) Headings(path string) ([]*heading.Heading, bool) {
	if result, found := r.headings[path]; found {
		return result, true
	}

	full := r.Absolute(path)

	if !system.FileExists(full) {
		return nil, false
	}

	result := heading.Parse(system.ReadFileUnsafe(full))
	r.headings[path] = result

	return result, true
}
