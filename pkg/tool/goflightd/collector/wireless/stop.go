package wireless

func (c *Collector) Stop() {
	close(c.stop)
}
