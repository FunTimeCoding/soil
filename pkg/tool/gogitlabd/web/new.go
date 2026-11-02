package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	r face.Reporter,
) *Server {
	g := registry.New()
	g.Register(
		palette.NewCommand(
			constant.BoardTitle,
			constant.BoardPath,
			webConstant.PaletteNavigate,
		),
	)

	return &Server{
		service:   v,
		inventory: i,
		reporter:  r,
		registry:  g,
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
