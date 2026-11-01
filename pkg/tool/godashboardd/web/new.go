package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
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
	v *service.Service,
	f face.Reporter,
) *Server {
	b := v.Board()
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

	for _, entry := range b.Entries() {
		r.Register(
			palette.NewCommand(entry.Label, entry.Link, webConstant.PaletteLink),
		)
	}

	return &Server{
		board:    b,
		service:  v,
		reporter: f,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeSentinel).
				WithStyle(
					join.Empty(webConstant.ComponentStyle, constant.InlineStyle),
				).
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
