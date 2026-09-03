package server

import (
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (s *Server) register(
	g *guard.Mux,
	pattern string,
	h http.Handler,
) {
	if s.openAuthentication {
		g.OpenMount(pattern, crossOriginMiddleware(s.authenticate(h)))

		return
	}

	g.TokenMount(pattern, crossOriginMiddleware(h))
}
