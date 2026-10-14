package watcher

import "github.com/funtimecoding/soil/pkg/errors"

func (w *Watcher) Close() {
	if w.notifier != nil {
		errors.PanicOnError(w.notifier.Close())
	}
}
