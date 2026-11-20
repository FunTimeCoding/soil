package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/board_result"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
)

func Tiles(b *board_result.Result) gomponents.Node {
	return tile_grid.New(
		[]gomponents.Node{
			count(b.Total, "Systems reporting"),
			count(b.Supported, "Supported release"),
			count(b.Tracked, "Fixes only in paid support"),
			count(b.Untracked, "No tracker for the release"),
			count(b.Upgradable, "An upgrade fixes something"),
			count(b.Stale, "Stale report"),
		},
	).Node()
}
