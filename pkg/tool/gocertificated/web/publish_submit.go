package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"net/http"
)

func (s *Server) publishSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	_, _, e := s.service.Publish()
	errors.PanicOnError(e)
	http.Redirect(w, r, constant.DashboardPath, http.StatusSeeOther)
}
