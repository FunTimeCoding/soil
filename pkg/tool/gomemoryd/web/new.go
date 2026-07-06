package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	theme "github.com/funtimecoding/soil/pkg/web/theme/constant"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *service.Service) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.MemoriesTitle,
			Path:     constant.MemoriesPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.ImpressionsTitle,
			Path:     constant.ImpressionsPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.SearchTitle,
			Path:     constant.SearchPath,
			Category: "navigate",
		},
		palette.Command{
			Label:      "Search memories",
			Path:       "/palette/memories",
			Category:   "search",
			SwapTarget: ".palette-body",
		},
	)

	return &Server{
		service:  s,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Cortex).
				WithStyle(inlineCSS).
				WithCommandPalette("/palette").
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
						constant.ImpressionsPath,
						constant.ImpressionsTitle,
					),
					navigation_item.New(
						constant.SearchPath,
						constant.SearchTitle,
					),
				),
		),
	}
}
