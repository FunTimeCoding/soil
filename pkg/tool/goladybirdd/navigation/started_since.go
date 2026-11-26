package navigation

func (w *Watcher) StartedSince(
	context string,
	mark int,
) bool {
	return w.wait(func() bool { return w.started[context] > mark }, w.window)
}
