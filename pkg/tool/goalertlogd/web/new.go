package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	v *service.Service,
	n face.EventNotifier,
) *Server {
	return &Server{
		service:  v,
		notifier: n,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeSentinel, favicon).
				WithStyle(constant.InlineStyle).
				WithLiveEndpoint(webConstant.LivePath).
				WithItems(
					navigation_item.New(
						constant.RecentPath,
						constant.RecentTitle,
					),
				),
		),
	}
}
