package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *store.Store,
	v *service.Service,
	authorization *client.Client,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.AuthoritiesTitle,
			Path:     constant.AuthoritiesPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.CertificatesTitle,
			Path:     constant.CertificatesPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.CreateAuthorityTitle,
			Path:     constant.CreateAuthorityPath,
			Category: web.PaletteAction,
		},
		palette.Command{
			Label:    constant.IssueCertificateTitle,
			Path:     constant.IssueCertificatePath,
			Category: web.PaletteAction,
		},
		palette.Command{
			Label:    constant.RootTitle,
			Path:     constant.RootPath,
			Category: web.PaletteAction,
		},
		palette.Command{
			Label:    web.SignOutTitle,
			Path:     web.SignOutPath,
			Category: web.PaletteAction,
		},
	)

	return &Server{
		store:         s,
		service:       v,
		authorization: authorization,
		registry:      registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeSentinel).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.AuthoritiesPath,
						constant.AuthoritiesTitle,
					),
					navigation_item.New(
						constant.CertificatesPath,
						constant.CertificatesTitle,
					),
				),
		),
	}
}
