package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(v *service.Service) *Server {
	return &Server{
		service: v,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeArchive, favicon).
				WithItems(
					navigation_item.New(
						constant.HeatmapPath,
						constant.HeatmapTitle,
					),
					navigation_item.New(
						constant.EventsPath,
						constant.EventsTitle,
					),
				),
		),
	}
}
