package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"github.com/funtimecoding/soil/pkg/web/layout"
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
	content := []gomponents.Node{html.H1(gomponents.Text(constant.UserTitle))}

	if message := form.ErrorText(r); message != "" {
		content = append(content, layout.Alert(message))
	}

	content = append(
		content,
		userTable(found),
		createUserForm(),
		passwordForm(),
	)
	s.view.RenderPage(w, constant.UserTitle, constant.UserPath, content...)
}
