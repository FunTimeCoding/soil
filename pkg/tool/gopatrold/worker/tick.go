package worker

import "context"

func (w *Worker) tick(x context.Context) []error {
	inspections, failures := w.round(x)

	for _, g := range w.gauges {
		g.Reset()
	}

	for _, i := range inspections {
		n := i.Instance
		w.reached.WithLabelValues(n).Set(flag(i.Reached))

		for _, a := range i.Arrays {
			w.required.WithLabelValues(n, a.Name).Set(float64(a.Required))
			w.active.WithLabelValues(n, a.Name).Set(float64(a.Active))
			w.failed.WithLabelValues(n, a.Name).Set(float64(a.Failed))
			w.rebuilding.WithLabelValues(n, a.Name).Set(flag(a.Rebuilding))
		}

		for _, d := range i.Drives {
			w.drive(n, d)
		}
	}

	return failures
}
