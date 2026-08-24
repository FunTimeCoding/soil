package worker

func (w *Worker) Stop() {
	close(w.stop)
}
