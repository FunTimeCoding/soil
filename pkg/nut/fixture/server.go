package fixture

import "net"

type Server struct {
	listener net.Listener
	answers  map[string]string
}
