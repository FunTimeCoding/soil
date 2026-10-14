package bluetooth

func (c *Collector) Close() {
	close(c.stop)
}
