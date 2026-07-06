package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"time"
)

func WaitForListen(
	port int,
	timeout time.Duration,
) bool {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		c, e := net.DialTimeout(
			constant.Transmission,
			fmt.Sprintf("localhost:%d", port),
			constant.Retry,
		)

		if e == nil {
			errors.PanicClose(c)

			return true
		}

		time.Sleep(constant.Retry)
	}

	return false
}
