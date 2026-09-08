package watcher

func (w *Watcher) Start() {
	w.mutex.Lock()

	if w.running {
		w.mutex.Unlock()

		return
	}

	w.running = true
	w.done = make(chan struct{})
	w.mutex.Unlock()
	w.selfIdentifier()
	subscriptions, e := w.store.All()

	if e != nil {
		w.reporter.CaptureException(e)
	}

	for _, v := range subscriptions {
		w.Index(v.RootIdentifier)
	}

	go w.listen()
	w.logger.Structured("watcher_start", "subscriptions", len(subscriptions))
}
