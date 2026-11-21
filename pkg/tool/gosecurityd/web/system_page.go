package web

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
)

func (s *Server) systemPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get(generativeConstant.ParameterName)
	m, e := minimum(r, 0)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	p, f := s.system.Packages(name, m)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	s.view.RenderPage(
		w,
		p.Summary.Name,
		constant.SystemsPath,
		html.H1(gomponents.Text(p.Summary.Name)),
		system_page.Facts(p.Summary),
		system_page.ScoreFilter(
			constant.SystemPath,
			url.Values{generativeConstant.ParameterName: {name}},
			m,
			constant.CriticalChips,
		),
		system_page.Kernel(p.Kernel),
		system_page.Packages("Upgrade fixes it", p.Upgrade, true),
		system_page.Packages(
			"Debian has a fix, not yet in the host's candidate",
			p.Pending,
			true,
		),
		system_page.Packages("Paid support only", p.Paid, false),
		system_page.Packages("No fix yet", p.Open, false),
	)
}
