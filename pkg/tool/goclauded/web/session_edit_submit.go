package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	"net/http"
)

func (s *Server) sessionEditSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(r.ParseForm())
	alias := r.FormValue(constant.Alias)
	description := r.FormValue(constant.Description)
	a := edit_session.New()
	a.Alias = &alias
	a.Description = &description
	errors.PanicOnError(s.service.EditSession(identifier, a))
	http.Redirect(
		w,
		r,
		fmt.Sprintf("/sessions/%s", identifier),
		http.StatusSeeOther,
	)
}
