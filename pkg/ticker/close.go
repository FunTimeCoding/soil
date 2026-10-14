package ticker

func (t *Ticker) Close() {
	t.ticker.Stop()
	close(t.done)
}
