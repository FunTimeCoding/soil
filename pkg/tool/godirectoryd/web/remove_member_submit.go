package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/url"
)

func (s *Server) removeMemberSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	_, e := s.service.RemoveMember(
		r.FormValue(constant.GroupField),
		r.FormValue(constant.AccountField),
	)

	if e != nil {
		form.RedirectFor(
			w,
			r,
			constant.GroupPath,
			constant.MemberForm,
			e.Error(),
			url.Values{
				constant.GroupField:   {r.FormValue(constant.GroupField)},
				constant.AccountField: {r.FormValue(constant.AccountField)},
			},
		)

		return
	}

	http.Redirect(w, r, constant.GroupPath, http.StatusSeeOther)
}
