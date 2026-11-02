package worker

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (w *Worker) run() {
	for _, p := range w.prepare {
		w.recovery.Run(func() { w.report(p()) })
	}

	v, e := fsnotify.NewWatcher()
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(v.Close()) }()

	if w.context.Err() != nil {
		return
	}

	for _, p := range w.paths {
		errors.PanicOnError(v.Add(p))
	}

	w.logger.Structured("watcher_start", "paths", w.paths)

	for {
		select {
		case <-w.context.Done():
			return
		case event, okay := <-v.Events:
			if !okay {
				return
			}

			w.recovery.Run(func() { w.report(w.handle(event)) })
		case f, okay := <-v.Errors:
			if !okay {
				return
			}

			panic(f)
		}
	}
}
