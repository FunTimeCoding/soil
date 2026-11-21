package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_result"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
)

func Tiles(r *upgrade_result.Result) gomponents.Node {
	return tile_grid.New(
		[]gomponents.Node{
			count(len(r.Groups), "Package upgrades"),
			count(r.Hosts, "Hosts with an upgrade"),
			count(r.Vulnerabilities, "Vulnerabilities an upgrade fixes"),
		},
	).Node()
}
