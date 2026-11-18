package fixture

import "net"

func (s *Server) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}
