package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	i *instance.Inventory,
	v *service.Service,
) *Server {
	return &Server{
		inventory: i,
		service:   v,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeForge, favicon).
				WithStyle(constant.InlineStyle).
				WithLiveEndpoint(webConstant.LivePath).
				WithItems(
					navigation_item.New(
						constant.FloorPath,
						constant.FloorTitle,
					),
				),
		),
	}
}
