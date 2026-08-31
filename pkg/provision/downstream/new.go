package downstream

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
	"strconv"
)

func New(address string) *Client {
	host, port, e := net.SplitHostPort(address)
	errors.PanicOnError(e)
	number, f := strconv.Atoi(port)
	errors.PanicOnError(f)

	return &Client{host: host, port: number}
}
