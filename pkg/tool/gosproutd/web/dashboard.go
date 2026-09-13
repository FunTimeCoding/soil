package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	r *http.Request,
) {
	modified := r.URL.Query().Get(constant.SortParameter) == constant.SortModified
	event := constant.Seeds
	table := s.seedTable()

	if modified {
		event = constant.SeedsRecent
		table = s.recentSeedTable()
	}

	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		html.H3(gomponents.Text("Seeds")),
		sortLinks(modified),
		html.Div(extended.StreamSwap(event), table),
	)
}
