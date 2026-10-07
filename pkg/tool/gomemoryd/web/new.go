package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *service.Service) *Server {
	r := registry.New()
	search := palette.NewCommand(
		"Search memories",
		"/palette/memories",
		web.PaletteSearch,
	)
	search.SwapTarget = ".palette-body"
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.MemoriesTitle,
			constant.MemoriesPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.RelationsTitle,
			constant.RelationsPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.ImpressionsTitle,
			constant.ImpressionsPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SearchTitle,
			web.SearchPath,
			web.PaletteNavigate,
		),
		search,
	)

	return &Server{
		service:  s,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeCortex).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.MemoriesPath,
						constant.MemoriesTitle,
					),
					navigation_item.New(
						constant.RelationsPath,
						constant.RelationsTitle,
					),
					navigation_item.New(
						constant.ImpressionsPath,
						constant.ImpressionsTitle,
					),
					navigation_item.New(
						constant.StatisticPath,
						constant.StatisticTitle,
					),
					navigation_item.New(web.SearchPath, constant.SearchTitle),
				),
		),
	}
}
