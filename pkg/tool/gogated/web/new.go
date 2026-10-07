package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *service.Service,
	authorization *client.Client,
	superUserMail string,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.ClientsTitle,
			constant.ClientsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CreateTitle,
			constant.CreatePath,
			webConstant.PaletteAction,
		),
		palette.NewCommand(
			constant.SessionsTitle,
			constant.SessionsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SignOutTitle,
			webConstant.SignOutPath,
			webConstant.PaletteAction,
		),
	)

	return &Server{
		service:       s,
		authorization: authorization,
		superUserMail: superUserMail,
		registry:      r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeAmethyst).
				WithStyle(constant.InlineCSS).
				WithCommandPalette("/palette").
				WithItems(
					navigation_item.New(
						constant.ClientsPath,
						constant.ClientsTitle,
					),
					navigation_item.New(
						constant.CreatePath,
						constant.CreateTitle,
					),
					navigation_item.New(
						constant.SessionsPath,
						constant.SessionsTitle,
					),
				),
		),
	}
}
