package chromium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
)

func Listening(
	host string,
	port int,
) bool {
	c, e := net.DialTimeout(
		constant.Transmission,
		fmt.Sprintf("%s:%d", host, port),
		constant.Retry,
	)

	if e != nil {
		return false
	}

	errors.PanicClose(c)

	return true
}
