package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/upgrade_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) upgrades(
	w http.ResponseWriter,
	r *http.Request,
) {
	show := r.URL.Query().Get(constant.ShowField)
	m, e := minimum(r, 0)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	u, f := s.system.Upgrades(show, m)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	s.view.RenderPage(
		w,
		constant.UpgradesTitle,
		constant.UpgradesPath,
		html.H1(gomponents.Text(constant.UpgradesTitle)),
		upgrade_page.Tiles(u),
		upgrade_page.Filter(show, u.Releases, m),
		upgrade_page.Table(u.Groups),
	)
}
