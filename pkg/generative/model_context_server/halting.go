package model_context_server

import "net/http"

func (s *Server) Halting() *http.Client {
	return s.pool.Halting()
}
