package model_context_server

import "net/http"

func (s *Server) Client() *http.Client {
	return s.pool.Client()
}
