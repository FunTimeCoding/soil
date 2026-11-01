package worker

import "github.com/funtimecoding/soil/pkg/errors/unavailable"

func (w *Worker) Poll() {
	v, e := w.service.Intents()

	if unavailable.Is(e) {
		w.logger.Plain("netbox status: %v", e)

		return
	}

	if e != nil {
		w.reporter.CaptureException(e)

		return
	}

	w.metric.SetIntents(v)
}
