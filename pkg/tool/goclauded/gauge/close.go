package gauge

func (g *Gauge) Close() {
	g.cancel()
	g.started.Do(func() { close(g.done) })
	<-g.done
}
