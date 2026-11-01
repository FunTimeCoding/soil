package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
)

func (s *Server) sessionDetail(
	w http.ResponseWriter,
	q *http.Request,
) {
	name := q.URL.Query().Get(constant.SessionParameter)

	if name == "" {
		http.Redirect(w, q, constant.SessionsPath, http.StatusSeeOther)

		return
	}

	s.view.RenderLivePage(
		w,
		name,
		constant.SessionsPath,
		fmt.Sprintf("%s=%s", constant.SessionParameter, url.QueryEscape(name)),
		html.H3(gomponents.Text(name)),
		html.Div(
			extended.StreamSwap(constant.DecisionEvent),
			session_page.Grid(s.service.Decisions(name)),
		),
	)
}
