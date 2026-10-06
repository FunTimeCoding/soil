package workspace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"path/filepath"
)

func New(
	diff bool,
	roots ...string,
) *Workspace {
	absolute := make([]string, 0, len(roots))

	for _, r := range roots {
		a, e := filepath.Abs(r)
		errors.PanicOnError(e)
		absolute = append(absolute, a)
	}

	return &Workspace{
		diff:     diff,
		roots:    absolute,
		before:   snapshot.Take(absolute...),
		overlay:  map[string][]byte{},
		original: map[string][]byte{},
	}
}
