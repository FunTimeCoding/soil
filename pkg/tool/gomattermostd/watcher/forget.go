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
	delete(w.since, root)
}
