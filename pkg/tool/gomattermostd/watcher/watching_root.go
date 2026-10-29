package watcher

func (w *Watcher) WatchingRoot(root string) bool {
	_, okay := w.rootOf(root)

	return okay
}
