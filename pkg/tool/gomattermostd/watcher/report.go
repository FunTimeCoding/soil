package watcher

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (w *Watcher) report(e error) {
	if classify.Reportable(e) {
		w.reporter.CaptureException(e)

		return
	}

	w.logger.Plain("watcher: %v", e)
}
