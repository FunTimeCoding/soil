package watcher

func (w *Watcher) Stop() {
	w.mutex.Lock()

	if !w.running {
		w.mutex.Unlock()

		return
	}

	w.running = false
	close(w.done)
	var pending []string

	for root, t := range w.timer {
		t.Stop()
		pending = append(pending, root)
	}

	w.mutex.Unlock()

	for _, root := range pending {
		w.Flush(root)
	}
}
