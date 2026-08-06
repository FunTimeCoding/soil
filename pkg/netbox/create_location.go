package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateLocation(
	name string,
	siteName string,
) (*location.Location, error) {
	s, e := c.SiteByName(siteName)

	if e != nil {
		return nil, e
	}

	q := netbox.NewWritableLocationRequest(
		name,
		slug(name),
		netbox.DeviceWithConfigContextRequestSite{
			Int32: &s.Identifier,
		},
	)
	result, _, f := c.client.DcimAPI.DcimLocationsCreate(
		c.context,
	).WritableLocationRequest(*q).Execute()

	if f != nil {
		return nil, f
	}

	return location.New(result), nil
}
