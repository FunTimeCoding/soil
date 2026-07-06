package network

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings"
	"net"
)

func SplitHostPort(s string) (string, int) {
	host, port, e := net.SplitHostPort(s)
	errors.PanicOnError(e)

	return host, strings.MustToInteger(port)
}
