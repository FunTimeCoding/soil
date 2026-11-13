package layer

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"path/filepath"
)

func New(
	parent *Layer,
	roots ...string,
) *Layer {
	result := &Layer{
		parent:   parent,
		written:  map[string][]byte{},
		removed:  map[string]bool{},
		original: map[string][]byte{},
		expected: map[string][]byte{},
		moves:    map[string]string{},
	}

	if parent != nil {
		result.roots = parent.roots
		result.before = parent.before

		return result
	}

	for _, r := range roots {
		a, e := filepath.Abs(r)
		errors.PanicOnError(e)
		result.roots = append(result.roots, a)
	}

	result.before = snapshot.Take(result.roots...)

	return result
}
