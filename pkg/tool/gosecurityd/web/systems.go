package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/selection_filter"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) systems(
	w http.ResponseWriter,
	r *http.Request,
) {
	selected, e := selection(r, 0)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	b, f := s.system.Board(selected)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	c, g := s.system.Facets(selected)

	if g != nil {
		s.view.Answer(w, r, s.reporter, g)

		return
	}

	s.view.RenderPage(
		w,
		constant.SystemsTitle,
		constant.SystemsPath,
		html.H1(gomponents.Text(constant.SystemsTitle)),
		system_page.Tiles(b),
		selection_filter.New(constant.SystemsPath, selected, c).Node(),
		system_page.Table(b.Rows),
		system_page.ScoreLine(s.score.Status()),
	)
}
