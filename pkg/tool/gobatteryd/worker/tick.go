package worker

func (w *Worker) tick() []error {
	p := w.read()
	w.reset()

	for server, reachable := range p.Reachable {
		value := 0.0

		if reachable {
			value = 1
		}

		w.reachable.WithLabelValues(server).Set(value)
	}

	for _, r := range p.Readings {
		w.record(r)
	}

	return p.Faults
}
