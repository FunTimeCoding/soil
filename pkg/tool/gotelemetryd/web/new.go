package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *store.Store) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.HeatmapTitle,
			constant.HeatmapPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.EventsTitle,
			constant.EventsPath,
			web.PaletteNavigate,
		),
	)

	return &Server{
		store:    s,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeArchive).
				WithCommandPalette(web.PalettePath).
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
