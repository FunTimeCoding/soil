package worker

func (w *Watcher) Start() {
	go w.recovery.Run(w.run)
}
