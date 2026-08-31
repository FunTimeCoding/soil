package chart

func (c *Chart) WithProjection() *Chart {
	c.projection = true

	return c
}
