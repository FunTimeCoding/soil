package worker

func (w *Worker) Close() {
	w.once.Do(func() { close(w.stop) })
}
