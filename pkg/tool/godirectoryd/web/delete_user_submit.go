package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/url"
)

func (s *Server) deleteUserSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())

	if e := s.service.DeleteUser(r.FormValue(constant.AccountField)); e != nil {
		form.Redirect(w, r, constant.UserPath, e.Error(), url.Values{})

		return
	}

	http.Redirect(w, r, constant.UserPath, http.StatusSeeOther)
}
