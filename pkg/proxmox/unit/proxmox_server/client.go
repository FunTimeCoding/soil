package proxmox_server

import (
	"github.com/funtimecoding/soil/pkg/proxmox"
	"net"
	"net/url"
	"strconv"
)

func (s *Server) Client() *proxmox.Client {
	s.t.Helper()
	l, e := url.Parse(s.server.URL)

	if e != nil {
		s.t.Fatalf("parse %s: %v", s.server.URL, e)
	}

	host, port, f := net.SplitHostPort(l.Host)

	if f != nil {
		s.t.Fatalf("split %s: %v", l.Host, f)
	}

	i, g := strconv.Atoi(port)

	if g != nil {
		s.t.Fatalf("port %s: %v", port, g)
	}

	result := proxmox.New(
		host,
		proxmox.WithToken("alfa@pve!test", "secret"),
		proxmox.WithPort(i),
		proxmox.WithUntrusted(),
	)
	s.t.Cleanup(result.Close)

	return result
}
