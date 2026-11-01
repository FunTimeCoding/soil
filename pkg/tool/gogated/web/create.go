package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/client_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) create(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.view.RenderPage(
		w,
		constant.CreateTitle,
		constant.CreatePath,
		html.H1(gomponents.Text(constant.CreateTitle)),
		client_page.CreateForm(r.URL.Query(), form.ErrorText(r)),
	)
}
