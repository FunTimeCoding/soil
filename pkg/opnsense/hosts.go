package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/host"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) Hosts(phrase string) ([]*host.Host, error) {
	rows, e := searchRows[response.Host](c, constant.HostSearch, phrase)

	if e != nil {
		return nil, e
	}

	return host.NewSlice(rows), nil
}
