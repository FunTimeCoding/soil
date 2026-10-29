package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	f face.Reporter,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.AuthoritiesTitle,
			constant.AuthoritiesPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CertificatesTitle,
			constant.CertificatesPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CreateAuthorityTitle,
			constant.CreateAuthorityPath,
			webConstant.PaletteAction,
		),
		palette.NewCommand(
			constant.IssueCertificateTitle,
			constant.IssueCertificatePath,
			webConstant.PaletteAction,
		),
		palette.NewCommand(
			constant.RootTitle,
			constant.RootPath,
			webConstant.PaletteAction,
		),
		palette.NewCommand(
			webConstant.SignOutTitle,
			webConstant.SignOutPath,
			webConstant.PaletteAction,
		),
	)

	return &Server{
		service:  v,
		reporter: f,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeSentinel).
				WithCommandPalette(webConstant.PalettePath).
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
