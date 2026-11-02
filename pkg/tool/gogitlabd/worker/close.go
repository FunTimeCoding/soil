package worker

func (w *Worker) Close() {
	w.stopped.Do(func() { close(w.stop) })
	w.started.Do(func() { close(w.done) })
	<-w.done
}
