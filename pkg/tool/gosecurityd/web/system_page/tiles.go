package system_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/board_result"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Tiles(b *board_result.Result) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			tile_grid.New(
				[]gomponents.Node{
					count(b.Total, "Systems reporting"),
					count(b.Supported, "Supported release"),
					count(b.Tracked, "Fixes only in paid support"),
					count(b.Untracked, "No tracker for the release"),
					count(b.Critical, "A vulnerability at 9.0 or higher"),
					count(b.Upgradable, "A package upgrade fixes something"),
					count(b.WithFix, "Package findings with a fix"),
					count(b.Stale, "Stale report"),
				},
			).Node(),
			tile_grid.New(
				[]gomponents.Node{
					count(
						b.Kernels[constant.KernelReboot],
						"Kernel: reboot fixes it",
					),
					count(
						b.Kernels[constant.KernelUpgrade],
						"Kernel: upgrade and reboot",
					),
					count(
						b.Kernels[constant.KernelStuck],
						"Kernel: no upgrade path",
					),
					count(
						b.Kernels[constant.KernelPaid],
						"Kernel: paid support only",
					),
					count(
						b.Kernels[constant.KernelCurrent],
						"Kernel: up to date",
					),
				},
			).Node(),
			html.P(
				html.Small(
					gomponents.Text(
						fmt.Sprintf(
							"%d more package findings have no fix yet, most of them not yet triaged by Debian.",
							b.NoFix,
						),
					),
				),
			),
		},
	)
}
