package tile_grid

import "maragu.dev/gomponents"

func New(tiles []gomponents.Node) *TileGrid {
	return &TileGrid{tiles: tiles}
}
