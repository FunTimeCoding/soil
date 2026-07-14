package server

import "time"

func (s *Server) WithWriteTimeout(d time.Duration) *Server {
	s.writeTimeout = d

	return s
}
