package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v face.Service,
	k *worker.Worker,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.FloorTitle,
			constant.FloorPath,
			web.PaletteNavigate,
		),
	)

	return &Server{
		service:  v,
		worker:   k,
		registry: r,
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
