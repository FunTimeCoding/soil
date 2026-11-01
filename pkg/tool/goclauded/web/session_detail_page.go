package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
	"net/url"
)

func (s *Server) sessionDetailPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	v, e := s.service.SessionView(q.Context(), q.PathValue(constant.Identifier))

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	p := session_page.New(v)
	s.view.RenderLivePage(
		w,
		p.Title(),
		constant.SessionsPath,
		join.Empty(
			subscription.Query(constant.Pulse),
			"&",
			url.Values{constant.EventSession: {v.Detail.Identifier}}.Encode(),
		),
		p.Detail(form.Failed(q, constant.CruiseAction)),
	)
}
