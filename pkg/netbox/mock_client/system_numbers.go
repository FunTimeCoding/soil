package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/system_number"

func (c *Client) SystemNumbers() ([]*system_number.Number, error) {
	return nil, c.Failure
}
