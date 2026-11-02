package watcher

func (w *Watcher) scanChanged() {
	for identifier, sourcePath := range w.changed {
		w.sync(identifier, sourcePath)
	}

	w.changed = make(map[string]string)
}
