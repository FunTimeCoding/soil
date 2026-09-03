package conversations

import "github.com/funtimecoding/soil/pkg/web/guard"

func (s *Server) Mount(g *guard.Mux) {
	g.Open("GET /conversations", s.page)
	g.Open("GET /conversations/sidebar", s.sidebar)
	g.Open("GET /conversations/{identifier}/edit", s.editForm)
	g.Open("POST /conversations/{identifier}/edit", s.editSubmit)
	g.Open("GET /conversations/{identifier}", s.panel)
}
