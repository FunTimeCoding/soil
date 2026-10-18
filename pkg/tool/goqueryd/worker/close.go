package worker

func (w *Worker) Close() {
	w.cancel()
	close(w.stop)
}
