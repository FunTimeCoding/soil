package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) outOfBand(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.mutex.Lock()
	s.count++
	count := s.count
	s.mutex.Unlock()
	s.view.RenderFragment(
		w,
		gomponents.Group(
			[]gomponents.Node{
				html.Span(
					gomponents.Textf(
						"received %s",
						r.FormValue(constant.TermField),
					),
				),
				html.Div(
					html.ID(constant.SummaryMark),
					extended.OutOfBand("true"),
					gomponents.Textf("sent %d", count),
				),
			},
		),
	)
}
