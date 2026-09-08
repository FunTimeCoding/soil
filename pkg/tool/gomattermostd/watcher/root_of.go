package watcher

func (w *Watcher) rootOf(identifier string) (string, bool) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	result, okay := w.index[identifier]

	return result, okay
}
