package locator

import (
	"net"
	"strconv"
)

func (l *Locator) Address() string {
	if l.port == 0 {
		return l.host
	}

	return net.JoinHostPort(l.host, strconv.Itoa(l.port))
}
