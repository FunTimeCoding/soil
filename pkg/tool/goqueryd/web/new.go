package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_cache"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *service.Service,
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
			constant.SearchTitle,
			webConstant.SearchPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CollectionsTitle,
			constant.CollectionsPath,
			webConstant.PaletteNavigate,
		),
	)

	return &Server{
		service:  s,
		registry: r,
		cache:    search_cache.New(10),
		reporter: f,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeSlate).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(webConstant.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						webConstant.SearchPath,
						constant.SearchTitle,
					),
					navigation_item.New(
						constant.CollectionsPath,
						constant.CollectionsTitle,
					),
				),
		),
	}
}
