package tile_grid

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (g *TileGrid) Node() gomponents.Node {
	grid := html.Div(
		html.Class(constant.TileGridClass),
		gomponents.Group(g.tiles),
	)

	if g.search == "" {
		return grid
	}

	return gomponents.Group{
		html.Input(
			html.Type(constant.SearchInputType),
			html.Placeholder(g.search),
			gomponents.Attr("oninput", "filterTiles(this)"),
		),
		grid,
		html.Script(gomponents.Raw(constant.TileFilterScript)),
	}
}
