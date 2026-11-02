package watch

func (w *Watch) Close() {
	w.clock.Close()
}
