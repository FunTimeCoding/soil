package navigation

func (w *Watcher) Settle(context string) bool {
	return w.wait(func() bool { return !w.open[context] }, w.limit)
}
