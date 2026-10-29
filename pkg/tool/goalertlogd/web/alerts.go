package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) alerts(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "missing name parameter", http.StatusBadRequest)

		return
	}

	records, e := s.service.Alerts(name)
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		name,
		"",
		html.H1(gomponents.Textf("Alert: %s", name)),
		html.P(gomponents.Textf("%d occurrences", len(records))),
		alertsTable(records),
	)
}
