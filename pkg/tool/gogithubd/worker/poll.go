package worker

func (w *Worker) Poll() {
	w.clock.Poll()
}
