package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/retry"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"time"
)

func WaitForListen(
	port int,
	timeout time.Duration,
) bool {
	_, e := retry.Until(
		timeout,
		constant.Retry,
		func() error {
			c, f := net.DialTimeout(
				constant.Transmission,
				fmt.Sprintf("localhost:%d", port),
				constant.Retry,
			)

			if f != nil {
				return f
			}

			errors.PanicClose(c)

			return nil
		},
	)

	return e == nil
}
