package worker

func (w *Worker) Close() {
	w.cancel()
	w.started.Do(func() { close(w.done) })
	<-w.done
}
