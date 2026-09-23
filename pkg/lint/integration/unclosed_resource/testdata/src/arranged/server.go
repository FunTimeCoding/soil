package example

import "net/http"

type Server struct {
	response *http.Response
}

func (s *Server) Close() {
	_ = s.response.Body.Close()
}

func (s *Server) Ping() {}

type Handle struct {
	response *http.Response
}

func (h *Handle) Close() error {
	return h.response.Body.Close()
}

func (h *Handle) Ping() {}
