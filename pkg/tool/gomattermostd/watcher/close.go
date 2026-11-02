package watcher

func (w *Watcher) Close() {
	w.mutex.Lock()

	if !w.running {
		w.mutex.Unlock()

		return
	}

	w.running = false
	w.mutex.Unlock()
	w.hold.Close()
}
