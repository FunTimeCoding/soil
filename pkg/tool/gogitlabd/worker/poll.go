package worker

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (w *Worker) Poll() {
	failures := w.refresh(w.context)

	if w.context.Err() != nil {
		return
	}

	for _, e := range failures {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)

			continue
		}

		w.logger.Structured("tick deferred", "error", e.Error())
	}

	w.refreshGauge()
}
