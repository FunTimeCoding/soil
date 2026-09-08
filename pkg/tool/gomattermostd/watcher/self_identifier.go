package watcher

func (w *Watcher) selfIdentifier() string {
	w.once.Do(
		func() {
			result, e := w.client.Me()

			if e != nil {
				w.reporter.CaptureException(e)

				return
			}

			w.self = result.Id
		},
	)

	return w.self
}
