package repository

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"go.yaml.in/yaml/v3"
	"maps"
	"os"
	"path/filepath"
	"slices"
)

func (r *Repository) Routes(directory string) ([]string, bool) {
	if paths, found := r.routes[directory]; found {
		return paths, true
	}

	if r.missing[directory] {
		return nil, false
	}

	p := r.Absolute(filepath.Join(directory, constant.RestSpecificationPath))

	if !system.FileExists(p) {
		r.missing[directory] = true

		return nil, false
	}

	b, e := os.ReadFile(p)
	errors.PanicOnError(e)
	var s specification
	errors.PanicOnError(yaml.Unmarshal(b, &s))
	paths := slices.Sorted(maps.Keys(s.Paths))
	r.routes[directory] = paths

	return paths, true
}
