package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
)

func Address() net.IP {
	addresses, e := net.InterfaceAddrs()
	errors.PanicOnError(e)

	for _, address := range addresses {
		if cast, success := address.(*net.IPNet); success &&
			!cast.IP.IsLoopback() {
			if cast.IP.To4() != nil {
				return cast.IP
			}
		}
	}

	return nil
}
