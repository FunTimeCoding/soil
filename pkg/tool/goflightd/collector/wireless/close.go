package wireless

func (c *Collector) Close() {
	close(c.stop)
}
