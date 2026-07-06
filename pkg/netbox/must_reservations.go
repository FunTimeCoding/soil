package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
)

func (c *Client) MustReservations() []*reservation.Reservation {
	result, e := c.Reservations()
	errors.PanicOnError(e)

	return result
}
