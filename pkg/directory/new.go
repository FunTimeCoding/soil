package directory

import "github.com/funtimecoding/soil/pkg/directory/constant"

func New(
	host string,
	base string,
) *Client {
	return &Client{host: host, port: constant.SecurePort, base: base}
}
