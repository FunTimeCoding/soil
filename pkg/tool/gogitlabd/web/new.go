package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	k *worker.Worker,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.BoardTitle,
			constant.BoardPath,
			webConstant.PaletteNavigate,
		),
	)

	return &Server{
		service:  v,
		worker:   k,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeTanuki).
				WithStyle(
					join.Empty(webConstant.ComponentStyle, constant.InlineStyle),
				).
				WithCommandPalette(webConstant.PalettePath).
				WithLiveEndpoint(webConstant.LivePath).
				WithItems(
					navigation_item.New(
						constant.BoardPath,
						constant.BoardTitle,
					),
				),
		),
	}
}
