package mock_collector

func (c *Collector) Fail(e error) {
	c.failure = e
}
