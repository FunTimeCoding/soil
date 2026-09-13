package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) branch(
	w http.ResponseWriter,
	r *http.Request,
) {
	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(
			w,
			html.Span(gomponents.Text(constant.FragmentContent)),
		)

		return
	}

	s.view.RenderFragment(w, html.Span(gomponents.Text(constant.PageContent)))
}
