package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/state"

func (c *Client) States(_ string) ([]*state.State, error) {
	return nil, nil
}
