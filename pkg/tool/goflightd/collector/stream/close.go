package stream

func (c *Collector) Close() {
	close(c.stop)
}
