package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"net/http"
	"time"
)

func (s *Server) cruiseFragment(
	w http.ResponseWriter,
	q *http.Request,
	identifier string,
	mode constant.Cruise,
	pace time.Duration,
	refused error,
) {
	if refused != nil {
		s.view.RenderFragment(
			w,
			session_page.RosterCruise(
				identifier,
				cruise_setting.New(identifier, mode, pace),
				refused.Error(),
			),
		)

		return
	}

	setting, e := s.message.CruiseSetting(identifier)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, session_page.RosterCruise(identifier, setting, ""))
}
