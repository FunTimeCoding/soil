package mechanic

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) tickCell() gomponents.Node {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return html.Span(gomponents.Textf("count %d", s.count))
}
