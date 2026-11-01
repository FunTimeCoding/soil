package tile

func (t *Tile) Terms(text string) *Tile {
	t.terms = text

	return t
}
