package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/remediation_page"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/selection_filter"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) remediationList(
	w http.ResponseWriter,
	r *http.Request,
) {
	selected, e := selection(r, constant.CriticalScore)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	result, f := s.remediation.List(selected)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	c, g := s.system.Facets(selected)

	if g != nil {
		s.view.Answer(w, r, s.reporter, g)

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
		selection_filter.New(constant.RemediationPath, selected, c).
			Fallback(constant.CriticalScore, constant.RemediationFloors).
			Node(),
		remediation_page.Table(open, "Nothing open at this score."),
		html.H2(gomponents.Text("Addressed by a decision")),
		remediation_page.Table(addressed, "No decision addresses an item yet."),
	)
}
