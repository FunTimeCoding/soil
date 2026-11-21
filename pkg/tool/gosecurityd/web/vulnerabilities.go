package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/vulnerability_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
)

func (s *Server) vulnerabilities(
	w http.ResponseWriter,
	r *http.Request,
) {
	m, e := minimum(r, constant.CriticalScore)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	groups, f := s.system.Groups(m)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	s.view.RenderPage(
		w,
		constant.VulnerabilitiesTitle,
		constant.VulnerabilitiesPath,
		html.H1(gomponents.Text(constant.VulnerabilitiesTitle)),
		vulnerability_page.Tiles(groups),
		system_page.ScoreFilter(
			constant.VulnerabilitiesPath,
			url.Values{},
			m,
			constant.FindingChips,
		),
		vulnerability_page.Table(groups),
	)
}
