package upstream_server

import "net/http"

func WithHeader(h http.Header) Option {
	return func(s *Server) {
		s.header = h
	}
}
