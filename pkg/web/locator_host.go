package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
	"strings"
)

func LocatorHost(s string) string {
	hostPort := LocatorHostPort(s)

	if strings.ContainsRune(hostPort, ':') {
		host, _, e := net.SplitHostPort(hostPort)

		if e != nil {
			if !strings.Contains(e.Error(), "missing port in address") {
				errors.PanicOnError(e)
			}
		}

		return host
	}

	return hostPort
}
