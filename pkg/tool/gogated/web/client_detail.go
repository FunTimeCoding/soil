package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) clientDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue("identifier")
	found, e := s.service.Client(identifier)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		"Client Detail",
		constant.ClientsPath,
		html.H1(gomponents.Textf("Client %s", truncate(identifier, 8))),
		clientDetailCard(found),
		html.Form(
			html.Method("post"),
			html.Action(fmt.Sprintf("/clients/%s/delete", identifier)),
			html.Button(
				html.Type("submit"),
				html.Class("secondary"),
				gomponents.Text("Delete Client"),
			),
		),
	)
}
