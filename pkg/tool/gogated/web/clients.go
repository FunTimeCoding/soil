package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/client_page"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) clients(
	w http.ResponseWriter,
	r *http.Request,
) {
	clients, e := s.service.ListClients()

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.ClientsTitle,
		constant.ClientsPath,
		html.H1(gomponents.Textf("Clients (%d)", len(clients))),
		load.New(constant.CreatePath, constant.CreateTitle).Button().Node(),
		client_page.Table(clients),
	)
}
