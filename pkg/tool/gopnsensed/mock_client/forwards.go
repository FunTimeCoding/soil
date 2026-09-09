package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/forward"

func (c *Client) Forwards(_ string) ([]*forward.Forward, error) {
	return nil, nil
}
