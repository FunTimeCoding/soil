package missing_host

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
)

func (c *Client) DeleteHost(identifier string) error {
	return not_found.New(constant.HostSubject, identifier)
}
