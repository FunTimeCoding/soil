package worker

func (w *Worker) Close() {
	w.cancel()
}
