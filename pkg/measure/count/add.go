package count

func (c *Count) Add(o *Count) {
	c.Code += o.Code
	c.Comment += o.Comment
	c.Blank += o.Blank
}
