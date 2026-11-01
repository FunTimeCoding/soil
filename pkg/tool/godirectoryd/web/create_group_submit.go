package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/url"
)

func (s *Server) createGroupSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	_, e := s.service.CreateGroup(r.FormValue(constant.NameField))

	if e != nil {
		form.RedirectFor(
			w,
			r,
			constant.GroupPath,
			constant.CreateGroupForm,
			e.Error(),
			url.Values{constant.NameField: {r.FormValue(constant.NameField)}},
		)

		return
	}

	http.Redirect(w, r, constant.GroupPath, http.StatusSeeOther)
}
