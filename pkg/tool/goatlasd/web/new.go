package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(v *service.Service) *Server {
	return &Server{
		service: v,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeTerritory, favicon).
				WithStyle(constant.InlineStyle).
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
