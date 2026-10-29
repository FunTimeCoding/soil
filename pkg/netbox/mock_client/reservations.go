package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/reservation"

func (c *Client) Reservations() ([]*reservation.Reservation, error) {
	return nil, c.Failure
}
