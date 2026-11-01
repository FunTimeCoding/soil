package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/seed_page"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	q *http.Request,
) {
	modified := q.URL.Query().Get(constant.SortParameter) ==
		constant.SortModified
	event := constant.Seeds
	table := seed_page.Priority(s.service.Seeds())

	if modified {
		event = constant.SeedsRecent
		table = seed_page.Recent(s.service.RecentSeeds())
	}

	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		html.H3(gomponents.Text(constant.SeedsTitle)),
		seed_page.Views(modified),
		html.Div(extended.StreamSwap(event), table),
	)
}
