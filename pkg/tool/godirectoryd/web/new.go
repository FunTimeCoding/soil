package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	authorization *client.Client,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.UserTitle,
			Path:     constant.UserPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.GroupTitle,
			Path:     constant.GroupPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    web.SignOutTitle,
			Path:     web.SignOutPath,
			Category: web.PaletteAction,
		},
	)

	return &Server{
		service:       v,
		authorization: authorization,
		registry:      registry,
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
