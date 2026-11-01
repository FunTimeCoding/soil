package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
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
	search := palette.NewCommand(
		"Search memories",
		"/palette/memories",
		webConstant.PaletteSearch,
	)
	search.SwapTarget = ".palette-body"
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.MemoriesTitle,
			constant.MemoriesPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.RelationsTitle,
			constant.RelationsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.ImpressionsTitle,
			constant.ImpressionsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SearchTitle,
			webConstant.SearchPath,
			webConstant.PaletteNavigate,
		),
		search,
	)

	return &Server{
		service:  s,
		reporter: f,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeCortex).
				WithStyle(
					join.Empty(webConstant.ComponentStyle, constant.InlineStyle),
				).
				WithCommandPalette(webConstant.PalettePath).
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
					navigation_item.New(
						webConstant.SearchPath,
						constant.SearchTitle,
					),
				),
		),
	}
}
