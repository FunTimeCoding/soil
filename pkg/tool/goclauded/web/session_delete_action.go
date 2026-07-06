package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) sessionDeleteAction(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(s.service.DeleteSession(identifier))
	w.Header().Set("HX-Redirect", constant.SessionsPath)
	w.WriteHeader(http.StatusOK)
}
