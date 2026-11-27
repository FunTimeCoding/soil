package chromium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"net"
)

func Listening(
	host string,
	port int,
) bool {
	c, e := net.DialTimeout(
		systemConstant.Transmission,
		fmt.Sprintf("%s:%d", host, port),
		systemConstant.Retry,
	)

	if e != nil {
		return false
	}

	errors.PanicClose(c)

	return true
}
