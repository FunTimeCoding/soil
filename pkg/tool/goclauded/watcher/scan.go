package watcher

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (w *Watcher) scan() {
	for _, e := range w.rescan(w.context) {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)
		}
	}
}
