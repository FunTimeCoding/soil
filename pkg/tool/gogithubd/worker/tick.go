package worker

func (w *Worker) tick() []error {
	latest, e := w.latest()

	if e != nil {
		return []error{e}
	}

	w.gauge.Reset()

	for k, r := range latest {
		w.gauge.WithLabelValues(
			w.owner,
			k.Repository,
			k.Workflow,
			k.Branch,
			r.Conclusion,
		).Set(1)
	}

	return nil
}
