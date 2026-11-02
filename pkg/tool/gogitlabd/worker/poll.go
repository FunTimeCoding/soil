package worker

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (w *Worker) Poll() {
	for _, e := range w.refresh() {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)

			continue
		}

		w.logger.Structured("tick deferred", "error", e.Error())
	}

	w.refreshGauge()
}
