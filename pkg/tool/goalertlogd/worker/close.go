package worker

func (w *Worker) Close() {
	w.closeOnce.Do(func() { close(w.stop) })
}
