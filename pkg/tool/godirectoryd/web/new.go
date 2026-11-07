package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(v *service.Service) *Server {
	return &Server{
		service: v,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemePhosphor, favicon).
				WithItems(
					navigation_item.New(constant.UserPath, constant.UserTitle),
					navigation_item.New(
						constant.GroupPath,
						constant.GroupTitle,
					),
				),
		),
	}
}
