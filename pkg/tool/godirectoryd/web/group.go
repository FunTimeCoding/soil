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

func (s *Server) group(
	w http.ResponseWriter,
	r *http.Request,
) {
	found, e := s.service.Groups()
	errors.PanicOnError(e)
	content := []gomponents.Node{html.H1(gomponents.Text(constant.GroupTitle))}

	if message := form.ErrorText(r); message != "" {
		content = append(content, layout.Alert(message))
	}

	content = append(
		content,
		groupTable(found),
		createGroupForm(),
		memberForm(),
	)
	s.view.RenderPage(w, constant.GroupTitle, constant.GroupPath, content...)
}
