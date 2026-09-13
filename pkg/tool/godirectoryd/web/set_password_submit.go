package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/url"
)

func (s *Server) setPasswordSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	account := r.FormValue(constant.AccountField)
	e := s.service.SetPassword(account, r.FormValue(constant.PasswordField))

	if e != nil {
		form.Redirect(
			w,
			r,
			constant.UserPath,
			e.Error(),
			url.Values{constant.AccountField: {account}},
		)

		return
	}

	http.Redirect(w, r, constant.UserPath, http.StatusSeeOther)
}
