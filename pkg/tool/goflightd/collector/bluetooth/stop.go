package bluetooth

func (c *Collector) Stop() {
	close(c.stop)
}
