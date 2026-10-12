package watch

func (w *Watch) Stop() {
	w.ticker.Stop()
}
