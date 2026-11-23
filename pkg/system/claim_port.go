package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net"
)

func ClaimPort() (int, net.Listener) {
	l, e := net.Listen(
		constant.Transmission,
		net.JoinHostPort(webConstant.Loopback, constant.AnyPort),
	)
	errors.PanicOnError(e)

	return l.Addr().(*net.TCPAddr).Port, l
}
