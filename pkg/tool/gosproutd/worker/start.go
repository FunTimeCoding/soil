package worker

func (w *Watcher) Start() {
	w.running.Add(1)
	go func() {
		defer w.running.Done()
		w.recovery.Run(w.run)
	}()
}
