package worker

import "context"

func (w *Worker) tick(x context.Context) []error {
	inspections, failures := w.round(x)

	for _, g := range w.gauges {
		g.Reset()
	}

	for _, i := range inspections {
		w.reached.WithLabelValues(i.Host).Set(flag(i.Reached))

		for _, a := range i.Arrays {
			w.required.WithLabelValues(i.Host, a.Name).Set(float64(a.Required))
			w.active.WithLabelValues(i.Host, a.Name).Set(float64(a.Active))
			w.failed.WithLabelValues(i.Host, a.Name).Set(float64(a.Failed))
			w.rebuilding.WithLabelValues(i.Host, a.Name).Set(flag(a.Rebuilding))
		}

		for _, d := range i.Drives {
			w.drive(i.Host, d)
		}
	}

	return failures
}
