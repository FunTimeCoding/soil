package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/request"

func (c *Client) SetHost(
	_ string,
	_ *request.Host,
) error {
	return nil
}
