package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	f face.Reporter,
) *Server {
	return &Server{
		service:  v,
		reporter: f,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeTyria, favicon).
				WithStyle(constant.InlineStyle).
				WithItems(
					navigation_item.New(constant.LogsPath, constant.LogsTitle),
					navigation_item.New(
						constant.ReportsPath,
						constant.ReportsTitle,
					),
					navigation_item.New(
						constant.RaidsPath,
						constant.RaidsTitle,
					),
					navigation_item.New(
						constant.PlayersPath,
						constant.PlayersTitle,
					),
				),
		),
	}
}
