package watcher

func (w *Watcher) watching(root string) bool {
	_, okay := w.rootOf(root)

	return okay
}
