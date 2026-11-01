package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
)

func (s *Server) deleteSession(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		s.service.DeleteAuthenticationSession(
			r.PathValue(constant.IdentifierParameter),
		),
	)
	http.Redirect(w, r, constant.SessionsPath, http.StatusFound)
}
