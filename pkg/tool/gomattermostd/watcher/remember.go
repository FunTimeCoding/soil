package watcher

func (w *Watcher) remember(
	identifier string,
	root string,
) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.index[identifier] = root
}
