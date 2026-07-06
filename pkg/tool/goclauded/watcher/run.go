package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (w *Watcher) run() {
	n, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	w.notifier = n
	w.addDirectories(n)
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	debounce := time.NewTimer(0)
	<-debounce.C

	for {
		select {
		case <-ticker.C:
			w.recovery.Run(
				func() {
					w.addDirectories(n)
					w.scan()
				},
			)
		case v, okay := <-n.Events:
			if !okay {
				return
			}

			if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) {
				continue
			}

			w.recordChange(v.Name)
			debounce.Reset(1 * time.Second)
		case <-debounce.C:
			if len(w.changed) > 0 {
				w.recovery.Run(w.scanChanged)
			}
		case _, okay := <-n.Errors:
			if !okay {
				return
			}
		}
	}
}
