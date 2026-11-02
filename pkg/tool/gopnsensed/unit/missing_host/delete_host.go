package missing_host

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	opnsenseConstant "github.com/funtimecoding/soil/pkg/opnsense/constant"
)

func (c *Client) DeleteHost(identifier string) error {
	c.Asked = append(c.Asked, identifier)

	return not_found.New(opnsenseConstant.HostSubject, identifier)
}
