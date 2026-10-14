package worker

func (w *Worker) Close() {
	close(w.stop)
}
