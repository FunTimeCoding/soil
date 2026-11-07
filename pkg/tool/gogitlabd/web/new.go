package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	r face.Reporter,
) *Server {
	return &Server{
		service:   v,
		inventory: i,
		reporter:  r,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeTanuki, favicon).
				WithStyle(constant.InlineStyle).
				WithLiveEndpoint(webConstant.LivePath).
				WithItems(
					navigation_item.New(
						constant.BoardPath,
						constant.BoardTitle,
					),
				),
		),
	}
}
