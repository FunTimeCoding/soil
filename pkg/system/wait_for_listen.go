package system

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/peer/retry"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"time"
)

func WaitForListen(
	port int,
	timeout time.Duration,
) bool {
	_, e := retry.Until(
		context.Background(),
		timeout,
		constant.Retry,
		func(x context.Context) error {
			d := net.Dialer{Timeout: constant.Retry}
			c, f := d.DialContext(
				x,
				constant.Transmission,
				fmt.Sprintf("localhost:%d", port),
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
