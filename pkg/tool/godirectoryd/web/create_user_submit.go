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
		form.RedirectFor(
			w,
			r,
			constant.UserPath,
			constant.CreateUserForm,
			e.Error(),
			url.Values{
				constant.AccountField: {account},
				constant.NameField:    {r.FormValue(constant.NameField)},
				constant.SurnameField: {r.FormValue(constant.SurnameField)},
				constant.MailField:    {r.FormValue(constant.MailField)},
			},
		)

		return
	}

	http.Redirect(w, r, constant.UserPath, http.StatusSeeOther)
}
