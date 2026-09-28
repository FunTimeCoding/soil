package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) clients(
	w http.ResponseWriter,
	r *http.Request,
) {
	clients := s.service.ListClients()
	s.view.RenderPage(
		w,
		constant.ClientsTitle,
		constant.ClientsPath,
		html.H1(gomponents.Textf("Clients (%d)", len(clients))),
		html.A(
			gomponents.Attr("href", constant.CreatePath),
			gomponents.Attr("role", "button"),
			gomponents.Text(constant.CreateTitle),
		),
		clientsTable(clients),
	)
}
