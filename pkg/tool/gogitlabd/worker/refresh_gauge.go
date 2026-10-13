package worker

func (w *Worker) refreshGauge() {
	w.gauge.Reset()

	for _, e := range w.Entries() {
		w.gauge.WithLabelValues(e.Instance, e.Project, e.Reference, e.Status).Set(
			1,
		)
	}
}
