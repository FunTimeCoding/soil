package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) clientDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue("identifier")
	clients := s.service.ListClients()
	var found *client.Client

	for _, c := range clients {
		if c.Identifier == identifier {
			found = c

			break
		}
	}

	if found == nil {
		http.NotFound(w, r)

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
