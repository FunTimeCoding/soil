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
	errors.PanicOnError(
		s.service.DeleteClient(r.PathValue(constant.IdentifierParameter)),
	)
	http.Redirect(w, r, constant.ClientsPath, http.StatusFound)
}
