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
	identifier := r.PathValue("identifier")
	errors.PanicOnError(s.service.DeleteAuthenticationSession(identifier))
	http.Redirect(w, r, constant.SessionsPath, http.StatusFound)
}
