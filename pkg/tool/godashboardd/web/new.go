package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(
	v *service.Service,
	b *board.Board,
	n face.EventNotifier,
	f face.Reporter,
) *Server {
	var links []*palette.Command

	for _, entry := range b.Entries() {
		links = append(
			links,
			palette.NewCommand(entry.Label, entry.Link, webConstant.PaletteLink),
		)
	}

	return &Server{
		board:    b,
		service:  v,
		notifier: n,
		reporter: f,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeSentinel, favicon).
				WithStyle(constant.InlineStyle).
				WithCommands(links...).
				WithLiveEndpoint(webConstant.LivePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.HeatmapPath,
						constant.HeatmapTitle,
					),
				).
				WithFooter(
					html.Script(gomponents.Raw(constant.BeaconScript)),
				),
		),
	}
}
