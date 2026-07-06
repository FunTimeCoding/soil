package watcher

import "github.com/funtimecoding/soil/pkg/errors"

func (w *Watcher) Stop() {
	if w.notifier != nil {
		errors.PanicOnError(w.notifier.Close())
	}
}
