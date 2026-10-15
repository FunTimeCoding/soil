package observe_transport

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"net/url"
	"strconv"
)

func host(u *url.URL) string {
	if u.Port() != "" {
		return u.Host
	}

	port := constant.SecurePort

	if u.Scheme == constant.Insecure {
		port = constant.InsecurePort
	}

	return net.JoinHostPort(u.Hostname(), strconv.Itoa(port))
}
