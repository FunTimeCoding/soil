package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_cache"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *service.Service) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SearchTitle,
			web.SearchPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CollectionsTitle,
			constant.CollectionsPath,
			web.PaletteNavigate,
		),
	)

	return &Server{
		service:  s,
		registry: r,
		cache:    search_cache.New(10),
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeSlate).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(web.SearchPath, constant.SearchTitle),
					navigation_item.New(
						constant.CollectionsPath,
						constant.CollectionsTitle,
					),
				),
		),
	}
}
