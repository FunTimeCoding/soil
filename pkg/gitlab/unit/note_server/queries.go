package note_server

import "net/url"

func (s *Server) Queries() []url.Values {
	s.lock.Lock()
	defer s.lock.Unlock()

	return append([]url.Values{}, s.queries...)
}
