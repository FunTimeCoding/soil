package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
)

func ClaimPort() (int, net.Listener) {
	l, e := net.Listen(constant.Transmission, ":0")
	errors.PanicOnError(e)

	return l.Addr().(*net.TCPAddr).Port, l
}
