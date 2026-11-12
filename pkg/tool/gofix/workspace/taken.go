package workspace

import "github.com/funtimecoding/soil/pkg/system"

func (w *Workspace) Taken(path string) bool {
	for _, to := range w.moves {
		if to == path {
			return true
		}
	}

	return system.FileExists(path)
}
