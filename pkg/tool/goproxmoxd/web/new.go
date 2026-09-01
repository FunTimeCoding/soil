package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v face.Service,
	k *worker.Worker,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.FloorTitle,
			Path:     constant.FloorPath,
			Category: web.PaletteNavigate,
		},
	)

	return &Server{
		service:  v,
		worker:   k,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeForge).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithLiveEndpoint(web.LivePath).
				WithItems(
					navigation_item.New(
						constant.FloorPath,
						constant.FloorTitle,
					),
				),
		),
	}
}
