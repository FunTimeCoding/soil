package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/component"
	"net/http"
)

func (s *Server) createRaid(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	identifier := s.service.CreateRaid(r.Form[constant.FileNameField])
	http.Redirect(w, r, component.RaidPath(identifier), http.StatusSeeOther)
}
