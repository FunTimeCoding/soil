package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *store.Store,
	authorization *client.Client,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     webConstant.RootPath,
			Category: webConstant.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.PlacementsTitle,
			Path:     constant.PlacementsPath,
			Category: webConstant.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.SightingsTitle,
			Path:     constant.SightingsPath,
			Category: webConstant.PaletteNavigate,
		},
		palette.Command{
			Label:    webConstant.SignOutTitle,
			Path:     webConstant.SignOutPath,
			Category: webConstant.PaletteAction,
		},
	)

	return &Server{
		store:         s,
		authorization: authorization,
		palette:       registry,
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
