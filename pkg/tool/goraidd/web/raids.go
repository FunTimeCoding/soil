package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) raids(
	w http.ResponseWriter,
	r *http.Request,
) {
	rows := s.store.Raids()
	s.view.RenderPage(
		w,
		constant.RaidsTitle,
		constant.RaidsPath,
		html.H1(gomponents.Textf("Raids (%d)", len(rows))),
		raidsTable(rows),
	)
}
