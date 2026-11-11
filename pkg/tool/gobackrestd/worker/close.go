package worker

func (w *Worker) Close() {
	w.clock.Close()
}
