package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/url"
)

func (s *Server) createUserSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	account := r.FormValue(constant.AccountField)
	_, e := s.service.CreateUser(
		account,
		r.FormValue(constant.NameField),
		r.FormValue(constant.SurnameField),
		r.FormValue(constant.MailField),
		r.FormValue(constant.PasswordField),
	)

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
