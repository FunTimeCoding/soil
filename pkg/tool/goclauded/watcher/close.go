package watcher

import "github.com/funtimecoding/soil/pkg/errors"

func (w *Watcher) Close() {
	w.cancel()
	w.started.Do(func() { close(w.done) })
	<-w.done

	if w.notifier != nil {
		errors.LogClose(w.notifier)
	}
}
