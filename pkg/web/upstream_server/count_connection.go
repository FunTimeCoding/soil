package upstream_server

import (
	"net"
	"net/http"
)

func (s *Server) countConnection(
	_ net.Conn,
	c http.ConnState,
) {
	if c != http.StateNew {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.opened++
}
