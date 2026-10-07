package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	authorization *client.Client,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.UserTitle,
			constant.UserPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.GroupTitle,
			constant.GroupPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(web.SignOutTitle, web.SignOutPath, web.PaletteAction),
	)

	return &Server{
		service:       v,
		authorization: authorization,
		registry:      r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemePhosphor).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(constant.UserPath, constant.UserTitle),
					navigation_item.New(
						constant.GroupPath,
						constant.GroupTitle,
					),
				),
		),
	}
}
