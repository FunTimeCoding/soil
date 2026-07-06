package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
)

func IsPortUnused(port int) bool {
	listener, e := net.Listen(
		constant.Transmission,
		fmt.Sprintf("localhost:%d", port),
	)

	if e != nil {
		return false
	}

	errors.LogClose(listener)

	return true
}
