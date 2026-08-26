package gohw

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"net"
)

func primaryInterface() (net.IP, net.IPMask, error) {
	interfaces, e := net.Interfaces()

	if e != nil {
		return nil, nil, e
	}

	for _, i := range interfaces {
		if i.Flags&net.FlagLoopback != 0 || i.Flags&net.FlagUp == 0 {
			continue
		}

		addresses, f := i.Addrs()

		if f != nil {
			continue
		}

		for _, a := range addresses {
			n, okay := a.(*net.IPNet)

			if !okay {
				continue
			}

			if n.IP.To4() != nil {
				return n.IP, n.Mask, nil
			}
		}
	}

	return nil, nil, not_found.Format("no non-loopback IPv4 interface")
}
