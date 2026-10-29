package worker

import "github.com/funtimecoding/soil/pkg/errors"

func (w *Worker) Poll() {
	latest, e := w.service.Latest()
	errors.PanicOnError(e)
	w.gauge.Reset()

	for k, r := range latest {
		w.gauge.WithLabelValues(
			w.service.Owner(),
			k.Repository,
			k.Workflow,
			k.Branch,
			r.Conclusion,
		).Set(1)
	}
}
