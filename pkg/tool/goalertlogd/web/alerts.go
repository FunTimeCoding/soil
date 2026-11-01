package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/alert_page"
	"net/http"
)

func (s *Server) alerts(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get(constant.NameParameter)

	if name == "" {
		http.Error(w, "missing name parameter", http.StatusBadRequest)

		return
	}

	records, e := s.service.Alerts(name)
	errors.PanicOnError(e)
	s.view.RenderPage(w, name, "", alert_page.Occurrences(name, records))
}
