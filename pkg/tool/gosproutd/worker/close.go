package worker

func (w *Watcher) Close() {
	w.closeOnce.Do(func() { close(w.stop) })
}
