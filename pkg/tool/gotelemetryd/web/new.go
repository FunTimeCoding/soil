package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *store.Store) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.HeatmapTitle,
			Path:     constant.HeatmapPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.EventsTitle,
			Path:     constant.EventsPath,
			Category: web.PaletteNavigate,
		},
	)

	return &Server{
		store:    s,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeArchive).
				WithCommandPalette("/palette").
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
