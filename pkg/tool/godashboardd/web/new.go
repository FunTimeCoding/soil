package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(
	b *board.Board,
	v *service.Service,
	c *store.Store,
	authorization *client.Client,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.HeatmapTitle,
			constant.HeatmapPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			webConstant.SignOutTitle,
			webConstant.SignOutPath,
			webConstant.PaletteAction,
		),
	)
	labels := map[string]bool{}

	for _, entry := range b.Entries() {
		labels[entry.Label] = true
		r.Register(
			palette.NewCommand(entry.Label, entry.Link, webConstant.PaletteLink),
		)
	}

	return &Server{
		board:         b,
		service:       v,
		store:         c,
		authorization: authorization,
		labels:        labels,
		registry:      r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeSentinel).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(webConstant.PalettePath).
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
