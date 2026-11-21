package flag_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
)

func Tiles(o *overview_result.Result) gomponents.Node {
	return tile_grid.New(
		[]gomponents.Node{
			count(o.Open, "Open"),
			count(o.Critical, "Critical, 9.0 or more"),
			count(o.DebianFix, "Debian fix available"),
			count(o.LocalAction, "Fix needs local action"),
		},
	).Node()
}
