package count

func (c *Count) Total() int {
	return c.Code + c.Comment + c.Blank
}
