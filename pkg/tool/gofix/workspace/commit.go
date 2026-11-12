package workspace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"maps"
	"os"
	"path/filepath"
	"slices"
)

func (w *Workspace) Commit() ([]string, error) {
	if len(w.order) == 0 {
		return nil, nil
	}

	if w.beforeCommit != nil {
		w.beforeCommit()
	}

	var written []string

	for _, path := range w.order {
		a, e := filepath.Abs(path)
		errors.PanicOnError(e)
		written = append(written, a)
	}

	var touched []string

	for _, root := range w.roots {
		if slices.ContainsFunc(
			written,
			func(path string) bool {
				return system.InsideDirectory(root, path)
			},
		) {
			touched = append(touched, root)
		}
	}

	if changed := w.before.Changed(touched...); len(changed) > 0 {
		return changed, nil
	}

	for _, path := range w.order {
		if e := os.WriteFile(path, w.overlay[path], 0644); e != nil {
			return nil, e
		}
	}

	for _, from := range slices.Sorted(maps.Keys(w.moves)) {
		to := w.moves[from]

		if system.FileExists(to) {
			return nil, fmt.Errorf("move %s: %s already exists", from, to)
		}

		if e := os.Rename(from, to); e != nil {
			return nil, e
		}
	}

	return nil, nil
}
