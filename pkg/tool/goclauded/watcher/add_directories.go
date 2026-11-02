package watcher

import "github.com/funtimecoding/soil/pkg/errors"

func (w *Watcher) addDirectories() {
	for _, d := range w.directories() {
		errors.PanicOnError(w.notifier.Add(d))
	}
}
