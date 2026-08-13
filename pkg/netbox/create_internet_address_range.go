package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateInternetAddressRange(
	start string,
	end string,
	status string,
	description string,
) (*internet_address_range.Range, error) {
	q := netbox.NewWritableIPRangeRequest(start, end)

	if status != "" {
		q.SetStatus(netbox.PatchedWritableIPRangeRequestStatus(status))
	}

	if description != "" {
		q.SetDescription(description)
	}

	result, _, e := c.client.IpamAPI.IpamIpRangesCreate(c.context).WritableIPRangeRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, e
	}

	return internet_address_range.New(result), nil
}
