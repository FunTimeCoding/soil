package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) CreatePrefix(
	_ string,
	_ *site.Site,
	_ string,
) (*prefix.Prefix, error) {
	return nil, nil
}
