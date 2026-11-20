package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) systems(
	w http.ResponseWriter,
	r *http.Request,
) {
	filter := r.URL.Query().Get(constant.ShowField)
	b, e := s.system.Board(filter)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		constant.SystemsTitle,
		constant.SystemsPath,
		html.H1(gomponents.Text(constant.SystemsTitle)),
		system_page.Tiles(b),
		system_page.Filter(filter),
		system_page.Table(b.Rows),
	)
}
