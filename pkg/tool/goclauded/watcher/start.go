package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (w *Watcher) Start() {
	w.started.Do(
		func() {
			n, e := fsnotify.NewWatcher()
			errors.PanicOnError(e)
			w.notifier = n
			w.addDirectories()
			go w.run()
		},
	)
}
