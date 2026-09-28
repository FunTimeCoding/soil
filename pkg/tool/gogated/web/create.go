package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/web/form"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var nodes []gomponents.Node
	nodes = append(nodes, html.H1(gomponents.Text(constant.CreateTitle)))

	if message := form.ErrorText(r); message != "" {
		nodes = append(nodes, layout.Alert(message))
	}

	q := r.URL.Query()
	nodes = append(
		nodes,
		createForm(q.Get("redirect_locators"), q.Get("scopes")),
	)
	s.view.RenderPage(
		w,
		constant.CreateTitle,
		constant.CreatePath,
		nodes...,
	)
}
