package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	c face.Forge,
	k *worker.Worker,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.BoardTitle,
			constant.BoardPath,
			web.PaletteNavigate,
		),
	)

	return &Server{
		client:   c,
		worker:   k,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeTanuki).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithLiveEndpoint(web.LivePath).
				WithItems(
					navigation_item.New(
						constant.BoardPath,
						constant.BoardTitle,
					),
				),
		),
	}
}
