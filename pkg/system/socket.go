package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
)

func Socket(address string) net.Conn {
	result, e := net.Dial("unix", address)
	errors.PanicOnError(e)

	return result
}
