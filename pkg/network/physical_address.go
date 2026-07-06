package network

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
)

func PhysicalAddress(s string) net.HardwareAddr {
	result, e := net.ParseMAC(s)
	errors.PanicOnError(e)

	return result
}
