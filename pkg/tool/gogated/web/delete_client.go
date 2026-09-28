package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
)

func (s *Server) deleteClient(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue("identifier")
	errors.PanicOnError(s.service.DeleteClient(identifier))
	http.Redirect(w, r, constant.ClientsPath, http.StatusFound)
}
