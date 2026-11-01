package tile_grid

func (g *TileGrid) Search(placeholder string) *TileGrid {
	g.search = placeholder

	return g
}
