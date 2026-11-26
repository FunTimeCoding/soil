package navigation

func (w *Watcher) Locator(context string) string {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.locators[context]
}
