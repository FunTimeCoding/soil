package stream

func (c *Collector) Stop() {
	close(c.stop)
}
