package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) sessionsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	page, e := strconv.Atoi(q.URL.Query().Get(constant.Page))

	if e != nil || page < 1 {
		page = 1
	}

	active := q.URL.Query().Get(constant.Active) == "true"
	all, f := s.service.EnrichedSessions(0, 0)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	sessions := all

	if active {
		sessions = nil

		for _, v := range all {
			if v.Active {
				sessions = append(sessions, v)
			}
		}
	}

	total := len(sessions)
	offset := min((page-1)*constant.PageSize, total)
	end := min(offset+constant.PageSize, total)
	s.view.RenderPage(
		w,
		constant.SessionsTitle,
		constant.SessionsPath,
		html.H3(gomponents.Text(constant.SessionsTitle)),
		session_page.List(
			[]*enriched_session.Session(sessions[offset:end]),
			active,
			page,
			total,
		),
	)
}
