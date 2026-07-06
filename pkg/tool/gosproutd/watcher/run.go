package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (w *Watcher) run() {
	w.recovery.Run(w.scan)
	n, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(n.Close()) }()
	w.addDirectories(n)
	sweep := time.NewTicker(10 * time.Minute)
	defer sweep.Stop()
	debounce := time.NewTimer(0)
	<-debounce.C
	var dirty bool

	for {
		select {
		case <-w.stop:
			return
		case <-sweep.C:
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

			if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) &&
				!v.Has(fsnotify.Remove) && !v.Has(fsnotify.Rename) {
				continue
			}

			dirty = true
			debounce.Reset(1 * time.Second)
		case <-debounce.C:
			if dirty {
				w.recovery.Run(
					func() {
						w.addDirectories(n)
						w.scan()
					},
				)
				dirty = false
			}
		case _, okay := <-n.Errors:
			if !okay {
				return
			}
		}
	}
}
