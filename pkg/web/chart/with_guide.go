package chart

func (c *Chart) WithGuide() *Chart {
	c.guide = true

	return c
}
