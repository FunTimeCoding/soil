package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox"
	"net"
	"net/url"
	"strconv"
)

func (u *Upstream) Client() *proxmox.Client {
	u.t.Helper()
	l, e := url.Parse(u.server.URL)

	if e != nil {
		u.t.Fatalf("parse %s: %v", u.server.URL, e)
	}

	host, port, f := net.SplitHostPort(l.Host)

	if f != nil {
		u.t.Fatalf("split %s: %v", l.Host, f)
	}

	p, g := strconv.Atoi(port)

	if g != nil {
		u.t.Fatalf("port %s: %v", port, g)
	}

	result := proxmox.New(
		host,
		proxmox.WithToken("alfa@pve!test", "secret"),
		proxmox.WithPort(p),
		proxmox.WithUntrusted(),
	)
	u.t.Cleanup(result.Close)

	return result
}
