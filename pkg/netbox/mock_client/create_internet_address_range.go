package mock_client

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
	if c.Failure != nil {
		return nil, c.Failure
	}

	c.record("CreateInternetAddressRange", start, end, status, description)
	v := netbox.NewIPRangeWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetStartAddress(start)
	v.SetEndAddress(end)

	return internet_address_range.New(v), nil
}
