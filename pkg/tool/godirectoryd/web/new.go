package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.GroupTitle,
			constant.GroupPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			webConstant.SignOutTitle,
			webConstant.SignOutPath,
			webConstant.PaletteAction,
		),
	)

	return &Server{
		service:       v,
		authorization: authorization,
		registry:      r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemePhosphor).
				WithCommandPalette(webConstant.PalettePath).
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
