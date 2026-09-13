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
		form.Redirect(w, r, constant.GroupPath, e.Error(), url.Values{})

		return
	}

	http.Redirect(w, r, constant.GroupPath, http.StatusSeeOther)
}
