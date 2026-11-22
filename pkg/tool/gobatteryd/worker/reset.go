package worker

func (w *Worker) reset() {
	for _, g := range w.measures {
		g.Reset()
	}

	w.status.Reset()
	w.information.Reset()
	w.attached.Reset()
	w.driver.Reset()
	w.reachable.Reset()
}
