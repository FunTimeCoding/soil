package watch

func (w *Watch) Start() {
	w.recovery.Run(w.Poll)
	w.ticker.Start()
}
