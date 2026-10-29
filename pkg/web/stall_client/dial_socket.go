package stall_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net"
)

func dialSocket(path string) func(
	context.Context,
	string,
	string,
) (net.Conn, error) {
	d := &net.Dialer{Timeout: constant.DialTimeout}

	return func(
		x context.Context,
		_ string,
		_ string,
	) (net.Conn, error) {
		return d.DialContext(x, "unix", path)
	}
}
