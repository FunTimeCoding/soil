package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/client_page"
	"net/http"
)

func (s *Server) clientDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	found, e := s.service.Client(r.PathValue(constant.IdentifierParameter))

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		"Client Detail",
		constant.ClientsPath,
		client_page.Detail(found),
	)
}
