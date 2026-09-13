package mechanic

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) counterCell() gomponents.Node {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return html.Span(gomponents.Textf("%d", s.count))
}
