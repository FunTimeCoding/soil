package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/selection_filter"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/upgrade_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) upgrades(
	w http.ResponseWriter,
	r *http.Request,
) {
	selected, e := selection(r, 0)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	release, kernel, f := upgradeFilter(r)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	u, g := s.system.Upgrades(selected, release, kernel)

	if g != nil {
		s.view.Answer(w, r, s.reporter, g)

		return
	}

	c, h := s.system.Facets(selected)

	if h != nil {
		s.view.Answer(w, r, s.reporter, h)

		return
	}

	s.view.RenderPage(
		w,
		constant.UpgradesTitle,
		constant.UpgradesPath,
		html.H1(gomponents.Text(constant.UpgradesTitle)),
		upgrade_page.Tiles(u),
		upgrade_page.Filter(selected, release, kernel, u.Releases),
		selection_filter.New(constant.UpgradesPath, selected, c).
			Keep(upgrade_page.Kept(release, kernel)).
			Node(),
		upgrade_page.Table(u.Groups),
	)
}
