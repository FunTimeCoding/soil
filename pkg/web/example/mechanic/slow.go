package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) slow(
	w http.ResponseWriter,
	_ *http.Request,
) {
	time.Sleep(constant.SlowWait)
	s.view.RenderFragment(w, html.Span(gomponents.Text("arrived")))
}
