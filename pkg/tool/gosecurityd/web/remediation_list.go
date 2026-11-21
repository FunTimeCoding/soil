package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/remediation_page"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
)

func (s *Server) remediationList(
	w http.ResponseWriter,
	r *http.Request,
) {
	m, e := minimum(r, constant.CriticalScore)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	result, f := s.remediation.List(m)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	var open, addressed []*remediation_item.Item

	for _, i := range result.Items {
		if i.Addressed {
			addressed = append(addressed, i)
		} else {
			open = append(open, i)
		}
	}

	s.view.RenderPage(
		w,
		constant.RemediationTitle,
		constant.RemediationPath,
		html.H1(gomponents.Text(constant.RemediationTitle)),
		remediation_page.Tiles(result),
		system_page.ScoreFilter(
			constant.RemediationPath,
			url.Values{},
			m,
			constant.FindingChips,
		),
		remediation_page.Table(open, "Nothing open at this score."),
		html.H2(gomponents.Text("Addressed by a decision")),
		remediation_page.Table(addressed, "No decision addresses an item yet."),
	)
}
