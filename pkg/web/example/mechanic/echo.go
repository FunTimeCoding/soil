package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) echo(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.view.RenderFragment(
		w,
		html.Span(
			gomponents.Textf(
				"term=%s scope=%s origin=%s",
				r.URL.Query().Get(constant.TermField),
				r.URL.Query().Get(constant.ScopeField),
				r.URL.Query().Get("origin"),
			),
		),
	)
}
