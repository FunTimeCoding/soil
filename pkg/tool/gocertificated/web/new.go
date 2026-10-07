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
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *store.Store,
	v *service.Service,
	authorization *client.Client,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.AuthoritiesTitle,
			constant.AuthoritiesPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CertificatesTitle,
			constant.CertificatesPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CreateAuthorityTitle,
			constant.CreateAuthorityPath,
			web.PaletteAction,
		),
		palette.NewCommand(
			constant.IssueCertificateTitle,
			constant.IssueCertificatePath,
			web.PaletteAction,
		),
		palette.NewCommand(
			constant.RootTitle,
			constant.RootPath,
			web.PaletteAction,
		),
		palette.NewCommand(web.SignOutTitle, web.SignOutPath, web.PaletteAction),
	)

	return &Server{
		store:         s,
		service:       v,
		authorization: authorization,
		registry:      r,
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
