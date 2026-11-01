package web

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) push(
	w http.ResponseWriter,
	name string,
	render func() (gomponents.Node, error),
) {
	n, e := render()

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	layout.PushEvent(w, name, n)
}
