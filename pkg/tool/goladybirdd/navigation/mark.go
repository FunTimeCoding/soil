package navigation

func (w *Watcher) Mark(context string) int {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.started[context]
}
