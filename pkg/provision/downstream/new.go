package downstream

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net"
	"strconv"
)

func New(address string) *Client {
	host, port, e := net.SplitHostPort(address)
	errors.PanicOnError(e)
	number, f := strconv.Atoi(port)
	errors.PanicOnError(f)

	return &Client{host: host, port: number, pool: stall_client.New().Build()}
}
