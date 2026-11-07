package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(
	s *service.Service,
	n face.EventNotifier,
	f face.Reporter,
) *Server {
	return &Server{
		service:  s,
		notifier: n,
		reporter: f,
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeSprout, favicon).
				WithStyle(constant.InlineStyle).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.SeedsTitle,
					),
					navigation_item.New(
						constant.SessionsPath,
						constant.SessionsTitle,
					),
					navigation_item.New(
						constant.CountersPath,
						constant.CountersTitle,
					),
				).
				WithScript(
					"https://cdn.jsdelivr.net/npm/sortablejs@1.15.6/Sortable.min.js",
				).
				WithLiveEndpoint(webConstant.LivePath).
				WithFooter(
					html.Script(gomponents.Raw(constant.SortableScript)),
				),
		),
	}
}
