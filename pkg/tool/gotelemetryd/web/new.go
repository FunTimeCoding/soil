package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(v *service.Service) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.HeatmapTitle,
			constant.HeatmapPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.EventsTitle,
			constant.EventsPath,
			webConstant.PaletteNavigate,
		),
	)

	return &Server{
		service:  v,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeArchive).
				WithCommandPalette(webConstant.PalettePath).
				WithItems(
					navigation_item.New(
						constant.HeatmapPath,
						constant.HeatmapTitle,
					),
					navigation_item.New(
						constant.EventsPath,
						constant.EventsTitle,
					),
				),
		),
	}
}
