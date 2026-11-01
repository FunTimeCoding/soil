package tile

func (t *Tile) Class(class string) *Tile {
	t.class = class

	return t
}
