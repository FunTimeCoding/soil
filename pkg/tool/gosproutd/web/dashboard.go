package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		html.H3(gomponents.Text("Seeds")),
		html.Div(
			gomponents.Attr("sse-swap", constant.Seeds),
			s.seedTable(),
		),
	)
}
