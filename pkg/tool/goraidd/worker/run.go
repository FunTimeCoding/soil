package worker

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (w *Worker) run() {
	w.recovery.Run(w.service.SyncLogCache)
	w.recovery.Run(w.service.EnrichExisting)
	w.recovery.Run(w.service.Cleanup)
	v, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(v.Close()) }()
	paths := w.service.WatchedPaths()

	if w.context.Err() != nil {
		return
	}

	for _, p := range paths {
		errors.PanicOnError(v.Add(p))
	}

	w.logger.Structured("watcher_start", "paths", paths)
	cleanup := time.NewTicker(24 * time.Hour)
	defer cleanup.Stop()

	for {
		select {
		case <-w.context.Done():
			return
		case <-cleanup.C:
			w.recovery.Run(w.service.Cleanup)
		case event, okay := <-v.Events:
			if !okay {
				return
			}

			w.recovery.Run(func() { w.service.HandleEvent(event) })
		case f, okay := <-v.Errors:
			if !okay {
				return
			}

			panic(f)
		}
	}
}
