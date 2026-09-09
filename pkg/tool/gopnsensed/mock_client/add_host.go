package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/request"

func (c *Client) AddHost(_ *request.Host) (string, error) {
	return "", nil
}
