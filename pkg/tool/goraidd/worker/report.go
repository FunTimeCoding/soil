package worker

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (w *Worker) report(failures []error) {
	for _, e := range failures {
		if classify.Reportable(e) {
			w.reporter.CaptureException(e)

			continue
		}

		w.logger.Structured("watch deferred", "error", e.Error())
	}
}
