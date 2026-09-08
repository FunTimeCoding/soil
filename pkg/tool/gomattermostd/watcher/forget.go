package watcher

func (w *Watcher) Forget(root string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for identifier, v := range w.index {
		if v == root {
			delete(w.index, identifier)
		}
	}

	delete(w.buffer, root)

	if t, okay := w.timer[root]; okay {
		t.Stop()
		delete(w.timer, root)
	}
}
