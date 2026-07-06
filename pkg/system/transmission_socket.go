package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
)

func TransmissionSocket(address string) net.Conn {
	result, e := net.Dial(constant.Transmission, address)
	errors.PanicOnError(e)

	return result
}
