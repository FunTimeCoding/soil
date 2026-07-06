package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
)

func (c *Client) Reservations() ([]*reservation.Reservation, error) {
	result, _, e := c.client.DcimAPI.DcimRackReservationsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return reservation.NewSlice(result.Results), nil
}
