package fixture

import "net"

func (s *Server) Host() string {
	return s.listener.Addr().(*net.TCPAddr).IP.String()
}
