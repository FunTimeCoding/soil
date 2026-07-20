package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_cache"
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
			Label:    constant.SearchTitle,
			Path:     constant.SearchPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.CollectionsTitle,
			Path:     constant.CollectionsPath,
			Category: "navigate",
		},
	)

	return &Server{
		service:  s,
		registry: registry,
		cache:    search_cache.New(10),
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Slate).
				WithStyle(constant.InlineStyle).
				WithCommandPalette("/palette").
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.SearchPath,
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
