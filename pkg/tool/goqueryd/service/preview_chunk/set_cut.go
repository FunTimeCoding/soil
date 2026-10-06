package preview_chunk

func (c *Chunk) SetCut(
	level int,
	lines []int,
) {
	c.CutChecked = true
	c.CutLevel = level
	c.CutLines = lines
}
