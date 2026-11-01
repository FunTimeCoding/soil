package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web/user_page"
	"github.com/funtimecoding/soil/pkg/web/component"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) user(
	w http.ResponseWriter,
	r *http.Request,
) {
	found, e := s.service.Users()
	errors.PanicOnError(e)
	values := r.URL.Query()
	s.view.RenderPage(
		w,
		constant.UserTitle,
		constant.UserPath,
		html.H1(gomponents.Text(constant.UserTitle)),
		component.Alert(form.PageError(r)),
		user_page.Table(found),
		user_page.CreateForm(values, form.Failed(r, constant.CreateUserForm)),
		user_page.PasswordForm(values, form.Failed(r, constant.PasswordForm)),
	)
}
