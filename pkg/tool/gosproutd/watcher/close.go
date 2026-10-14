package watcher

func (w *Watcher) Close() {
	close(w.stop)
}
