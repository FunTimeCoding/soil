package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *store.Store,
	authorization *client.Client,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			webConstant.RootPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.PlacementsTitle,
			constant.PlacementsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SightingsTitle,
			constant.SightingsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			webConstant.SignOutTitle,
			webConstant.SignOutPath,
			webConstant.PaletteAction,
		),
	)

	return &Server{
		store:         s,
		authorization: authorization,
		palette:       r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeTerritory).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(webConstant.PalettePath).
				WithItems(
					navigation_item.New(
						webConstant.RootPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.PlacementsPath,
						constant.PlacementsTitle,
					),
					navigation_item.New(
						constant.SightingsPath,
						constant.SightingsTitle,
					),
				),
		),
	}
}
