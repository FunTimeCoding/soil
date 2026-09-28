package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *service.Service,
	authorization *client.Client,
	superUserMail string,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.ClientsTitle,
			Path:     constant.ClientsPath,
			Category: webConstant.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.CreateTitle,
			Path:     constant.CreatePath,
			Category: webConstant.PaletteAction,
		},
		palette.Command{
			Label:    constant.SessionsTitle,
			Path:     constant.SessionsPath,
			Category: webConstant.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.SignOutTitle,
			Path:     webConstant.SignOutPath,
			Category: webConstant.PaletteAction,
		},
	)

	return &Server{
		service:       s,
		authorization: authorization,
		superUserMail: superUserMail,
		registry:      registry,
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
