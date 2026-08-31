package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(k *worker.Worker) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.PlateTitle,
			Path:     constant.PlatePath,
			Category: web.PaletteNavigate,
		},
	)

	return &Server{
		worker:   k,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeAtlas).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithLiveEndpoint(web.LivePath).
				WithItems(
					navigation_item.New(
						constant.PlatePath,
						constant.PlateTitle,
					),
				),
		),
	}
}
