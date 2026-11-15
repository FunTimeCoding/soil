package issuer_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"net/http"
)

func (s *Server) Mux() *http.ServeMux {
	s.t.Helper()
	result := http.NewServeMux()
	g := guard.New(result, []string{"alfa"})
	c := s.Client(memory.New())
	s.t.Cleanup(c.Close)
	c.Mount(g)
	g.Session("GET /target", web_tester.Serve)

	return result
}
