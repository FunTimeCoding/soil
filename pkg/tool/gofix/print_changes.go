package gofix

import "github.com/funtimecoding/soil/pkg/tool/gofix/workspace"

func printChanges(w *workspace.Workspace) {
	for _, path := range w.Changes() {
		printDiff(path, w.Original(path), w.Current(path))
	}
}
