package upstream_server

import "net/http"

func (s *Server) Last() *http.Request {
	s.t.Helper()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.requests) == 0 {
		s.t.Fatalf("no request reached %s", s.URL)
	}

	return s.requests[len(s.requests)-1]
}
