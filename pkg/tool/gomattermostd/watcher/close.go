package watcher

func (w *Watcher) Close() {
	w.mutex.Lock()

	if !w.running {
		w.mutex.Unlock()

		return
	}

	w.running = false
	var pending []string

	for root, t := range w.timer {
		t.Stop()
		pending = append(pending, root)
	}

	w.mutex.Unlock()
	w.hold.Close()

	for _, root := range pending {
		w.Flush(root)
	}
}
