package opnsense

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) ReconfigureDnsmasq() error {
	b, e := c.basic.Post(constant.DnsmasqReconfigure, struct{}{})

	if e != nil {
		return e
	}

	var out response.Status

	if f := json.Unmarshal(b, &out); f != nil {
		return f
	}

	if out.Status != constant.OkayStatus {
		return unexpected.Format(
			"unexpected dnsmasq reconfigure: %s",
			out.Status,
		)
	}

	return nil
}
