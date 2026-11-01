package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
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
			constant.LogsTitle,
			constant.LogsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.ReportsTitle,
			constant.ReportsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.RaidsTitle,
			constant.RaidsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.PlayersTitle,
			constant.PlayersPath,
			webConstant.PaletteNavigate,
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
				WithTheme(webConstant.ThemeTyria).
				WithStyle(
					join.Empty(webConstant.ComponentStyle, constant.InlineStyle),
				).
				WithCommandPalette(webConstant.PalettePath).
				WithItems(
					navigation_item.New(constant.LogsPath, constant.LogsTitle),
					navigation_item.New(
						constant.ReportsPath,
						constant.ReportsTitle,
					),
					navigation_item.New(
						constant.RaidsPath,
						constant.RaidsTitle,
					),
					navigation_item.New(
						constant.PlayersPath,
						constant.PlayersTitle,
					),
				),
		),
	}
}
