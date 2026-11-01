package tile

import "maragu.dev/gomponents"

func New(content ...gomponents.Node) *Tile {
	return &Tile{content: content}
}
