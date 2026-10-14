package watch

func (w *Watch) Close() {
	w.ticker.Close()
}
