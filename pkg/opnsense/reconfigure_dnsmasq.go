package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) ReconfigureDnsmasq() error {
	var out response.Status

	if e := c.basic.Post(
		constant.DnsmasqReconfigure,
		struct{}{},
		&out,
	); e != nil {
		return e
	}

	if out.Status != constant.OkayStatus {
		return unexpected.Format(
			"unexpected dnsmasq reconfigure: %s",
			out.Status,
		)
	}

	return nil
}
