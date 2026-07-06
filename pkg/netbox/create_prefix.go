package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePrefix(
	cidr string,
	s *site.Site,
	description string,
) (*prefix.Prefix, error) {
	q := netbox.NewWritablePrefixRequest(cidr)
	q.SetStatus(netbox.PATCHEDWRITABLEPREFIXREQUESTSTATUS_ACTIVE)

	if s != nil {
		q.SetScopeType("dcim.site")
		q.SetScopeId(s.Identifier)
	}

	if description != "" {
		q.SetDescription(description)
	}

	result, _, e := c.client.IpamAPI.IpamPrefixesCreate(
		c.context,
	).WritablePrefixRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return prefix.New(result), nil
}
